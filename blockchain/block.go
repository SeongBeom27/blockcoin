package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/baaami/blockcoin/db"
	"github.com/baaami/blockcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	// block 해싱
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	// block을 db에 저장
	block.persist()
	return &block
}
