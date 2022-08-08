package main

import (
	"fmt"
	"os"
)

func usage(){
	fmt.Printf("Welcome to block coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API\n")	
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start explorer")
	case "rest":	
		fmt.Println("Start rest")
	default:
		usage()
	}
}