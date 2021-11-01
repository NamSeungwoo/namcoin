package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/NamSeungwoo/namcoin/utils"
)

const (
	fileName string = "namcoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("namcoin.wallet")
	return !os.IsNotExist(err)
}

func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key) // the byte isn't need to be a string based HEX
	utils.HandleErr(err)
	err = ioutil.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)

}

func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsBytes, err := ioutil.ReadFile(fileName)
	utils.HandleErr(err)
	key, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return
}

func aFromK(key *ecdsa.PrivateKey) string {
	z := append(key.X.Bytes(), key.Y.Bytes()...)
	return fmt.Sprintf("%x", z)
}

func sign(payload string, w *wallet) string {
	payloadAsB, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsB)
	utils.HandleErr(err)
	signature := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signature)
}

// to sign a message is that create a signature for message but not encryption of message
// need both message and signature for verification

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	firstHalfBytes := bytes[:len(bytes)/2]
	secondHalfBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(secondHalfBytes)
	return &bigA, &bigB, nil
}

func verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	payloadBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&publicKey, payloadBytes, r, s)
	return ok
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = aFromK(w.privateKey)
	}
	return w
}

// const (
// 	signature     string = "486386b9375e1fde8eb026f295a337e61effe99c9922b6b33c3b41d63dd7b6608b9f669ced700e068b533d4269d584742d0e88a33cf7469380eb3966ad5cffb6"
// 	privateKey    string = "30770201010420351a7ad93c9a02c80880cf794b1da0cd02e51c1fc6577fdd4bf29cdd3f139b0da00a06082a8648ce3d030107a14403420004308e6479a40130b52b9e0f1629b33a2f0624a1a033317c6547d7bfa9effff01e5d4fe47138bb1dbc20763450dabd4c5a1f3081ad96d692e48ad1ffbc34a379d8"
// 	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
// )

func Start() {

	// privBytes, err := hex.DecodeString(privateKey)
	// utils.HandleErr(err)

	// private, err := x509.ParseECPrivateKey(privBytes)
	// utils.HandleErr(err)

	// bytes, err := hex.DecodeString(signature)

	// rBytes := bytes[:len(bytes)/2]
	// sBytes := bytes[len(bytes)/2:]

	// var bigR, bigS = big.Int{}, big.Int{}

	// bigR.SetBytes(rBytes)
	// bigS.SetBytes(sBytes)

	// hashBytes, err := hex.DecodeString(hashedMessage)

	// ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)

	// fmt.Println(ok)

}
