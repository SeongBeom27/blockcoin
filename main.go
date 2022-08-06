package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/baaami/blockcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templ.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal : error handling
	log.Fatal(http.ListenAndServe(port, nil))
}