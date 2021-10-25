package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("namcoin.wallet")
	return !os.IsNotExist(err)
}

func Wallet() *wallet {
	if w == nil {
		if hasWalletFile() {

		}
	}
	return w
}

// const (
// 	signature     string = "486386b9375e1fde8eb026f295a337e61effe99c9922b6b33c3b41d63dd7b6608b9f669ced700e068b533d4269d584742d0e88a33cf7469380eb3966ad5cffb6"
// 	privateKey    string = "30770201010420351a7ad93c9a02c80880cf794b1da0cd02e51c1fc6577fdd4bf29cdd3f139b0da00a06082a8648ce3d030107a14403420004308e6479a40130b52b9e0f1629b33a2f0624a1a033317c6547d7bfa9effff01e5d4fe47138bb1dbc20763450dabd4c5a1f3081ad96d692e48ad1ffbc34a379d8"
// 	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
// )

// func Start() {

// 	privBytes, err := hex.DecodeString(privateKey)
// 	utils.HandleErr(err)

// 	private, err := x509.ParseECPrivateKey(privBytes)
// 	utils.HandleErr(err)

// 	sigBytes, err := hex.DecodeString(signature)

// 	rBytes := sigBytes[:len(sigBytes)/2]
// 	sBytes := sigBytes[len(sigBytes)/2:]

// 	var bigR, bigS = big.Int{}, big.Int{}

// 	bigR.SetBytes(rBytes)
// 	bigS.SetBytes(sBytes)

// 	hashBytes, err := hex.DecodeString(hashedMessage)

// 	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)

// 	fmt.Println(ok)

// }
