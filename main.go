package main

import (
	"github.com/NAM/namcoin/cli"
	"github.com/NAM/namcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
