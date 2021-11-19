package main

import (
	"github.com/NamSeungwoo/namcoin/cli"
	"github.com/NamSeungwoo/namcoin/db"
)

func main() {

	// hash := sha256.Sum256([]byte("hello"))
	// fmt.Printf("%x\n", hash)
	defer db.Close()
	cli.Start()
	//wallet.Wallet()
}
