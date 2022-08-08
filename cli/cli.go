package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/baaami/blockcoin/explorer"
	"github.com/baaami/blockcoin/rest"
)

func usage(){
	fmt.Printf("Welcome to block coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:		Set the PORT of the server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n")	
	os.Exit(1)
}

func Start() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}