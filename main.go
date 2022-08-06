package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	// `json:"url"` : struct field tags 사용
	URL			string 	`json:"url"`
	Method 		string 	`json:"method"`
	Description string 	`json:"description"`
	// omitempty : data가 비어있지 않은 경우에만 출력
	Payload		string	`json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{			
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: "/blokcs",
			Method: "POST",
			Description: "See Documentation",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}