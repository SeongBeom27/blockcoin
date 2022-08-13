package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/baaami/blockcoin/db"
	"github.com/baaami/blockcoin/utils"
)

const difficulty int = 2

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

var ErrNotFound = errors.New("Block not found ")

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	// insert "00"
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := utils.Hash(b)
		fmt.Printf("Target:%s\nHash:%s\nNonce:%d\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Timestamp = int(time.Now().Unix())
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := Block{
		Hash:         "",
		PrevHash:     prevHash,
		Height:       height,
		Difficulty:   Blockchain().difficulty(),
		Nonce:        0,
		Transactions: []*Tx{makeCoinbaseTx("baami")},
	}
	// 작업 증명
	block.mine()
	// block을 db에 저장
	block.persist()
	return &block
}
