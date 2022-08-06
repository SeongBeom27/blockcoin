package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/baaami/blockcoin/blockchain"
	"github.com/baaami/blockcoin/utils"
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

type AddBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{			
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: "/blocks",
			Method: "POST",
			Description: "See Documentation",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)

	// input writer -> data Marshaling -> transfer json data to writer
	json.NewEncoder(rw).Encode(data)

}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		// {"data": "my block data"}

		// json into struct

		// 1. data 구조 확인
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}