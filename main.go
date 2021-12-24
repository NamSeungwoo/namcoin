package main

import (
	"github.com/NamSeungwoo/namcoin/cli"
	"github.com/NamSeungwoo/namcoin/db"
)

func main() {

	defer db.Close()
	cli.Start()

}
