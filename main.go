package main

import (
	"github.com/baaami/blockcoin/cli"
	"github.com/baaami/blockcoin/db"
)

func main() {
	defer db.DB().Close()
	cli.Start()
}
