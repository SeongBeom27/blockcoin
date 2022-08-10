package blockchain

import (
	"sync"

	"github.com/baaami/blockcoin/db"
	"github.com/baaami/blockcoin/utils"
)

type blockchain struct {
	// 1. 새로운 블록을 만들기 위해서 마지막 해쉬를 알아야함
	NewestHash string `json:"newestHash"`
	// 2. Height를 알아야함
	Height int `json:"height"`
}

// sigleton
var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// singleton 인스턴스 획득
func Blockchain() *blockchain {
	if b == nil {
		// 처음 초기화하는 부분은 반드시 1번만 수행되도록
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}
	return b
}
