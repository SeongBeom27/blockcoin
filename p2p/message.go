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

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		err := json.Unmarshal(m.Payload, &payload)
		utils.HandleErr(err)

		fmt.Println(payload)
	}
	// fmt.Printf("Peer: %s, Sent a message with kind of: %d", p.key, m.Kind)
}
