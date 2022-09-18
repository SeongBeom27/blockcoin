package p2p

import (
	"encoding/json"

	"github.com/baaami/blockcoin/blockchain"
	"github.com/baaami/blockcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

func (m *Message) addPayload(p interface{}) {
	b, err := json.Marshal(p)
	utils.HandleErr(err)
	m.Payload = b
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}
	m.addPayload(payload)
	mJson, err := json.Marshal(m)
	utils.HandleErr(err)
	return mJson
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
