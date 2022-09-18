package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/baaami/blockcoin/blockchain"
	"github.com/baaami/blockcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	fmt.Printf("Sending newest block to %s\n", p.key)
	
	// 마지막 해시값
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)

	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

type Message struct {
	// Message 종류 유형
	Kind    MessageKind
	Payload []byte
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Received the newest block form %s\n", p.key)

		var payload blockchain.Block
		err := json.Unmarshal(m.Payload, &payload)
		utils.HandleErr(err)

		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)

		// Execution 3000 port this code
		if payload.Height >= b.Height {
			// 3000번 포트가 블록이 부족하기 때문에 받아야함
			// request all the blocks from 4000
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requestAllBlocks(p)
		} else {
			// 3000번 포트가 블록이 많기 때문에 보내줘야함
			// send 4000 our block
			fmt.Printf("Sending newest block to %s\n", p.key)
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks.\n", p.key)
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Received all the blocks %s wants all the blocks.\n", p.key)
		var payload []blockchain.Block
		err := json.Unmarshal(m.Payload, &payload)
		utils.HandleErr(err)
	}
}
