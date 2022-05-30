package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "3077020101042019c4ed17f7ccbdb6f317623e4741b6cb6dddb6caff39eba47e9565e31faa1a87a00a06082a8648ce3d030107a1440342000421e3fc0ba79e68e63559940a52f34425c449cde202f065e7b80bb7fab5fd4a11f85be4b261a1ee2916d2082083d8bbc55b14c4948269a34e0162964749d3a5a9"
	testPayload string = "00fad70befde16d2a9630e7d00cc0c463c42700769f14440ee12920b02fc8630"
	testSig     string = "47ae5f68c41eeeea6e66d941040504a3ed3c5c356a8c0003e08972d7e20fcf06f4a747f0d9d705f3b96cf866833ea26fb666f190193fb29aee8160bfe65068b5"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}

	tests := []test{
		{testPayload, true},
		{"00fad70befde16d2a9630e7d00cc0c463c42700769f14440ee12920b02fc8631", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != !tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}

}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
