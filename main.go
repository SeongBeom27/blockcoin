package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	// Fprint : writer에 출력험
	fmt.Fprint(w, "Hello from Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal : error handling
	log.Fatal(http.ListenAndServe(port, nil))
}