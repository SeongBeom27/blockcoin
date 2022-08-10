package main

import (
	"github.com/baaami/blockcoin/blockchain"
	"github.com/baaami/blockcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
