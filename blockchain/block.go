package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/baaami/blockcoin/db"
	"github.com/baaami/blockcoin/utils"
)

const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
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
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
		Difficulty: difficulty,
		Nonce: 		0,
	}
	// 작업 증명
	block.mine()
	// block을 db에 저장
	block.persist()
	return &block
}
