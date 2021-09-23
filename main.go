package main

import (
	"github.com/NAM/namcoin/cli"
	"github.com/NAM/namcoin/db"
)

func main() {

	// hash := sha256.Sum256([]byte("hello"))
	// fmt.Printf("%x\n", hash)
	defer db.Close()
	cli.Start()
}
