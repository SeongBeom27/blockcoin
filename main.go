package main

import (
	"github.com/baaami/blockcoin/explorer"
	"github.com/baaami/blockcoin/rest"
)

func main() {
	go rest.Start(4000)
	explorer.Start(3000)
}