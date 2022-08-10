package blockchain

import (
	"fmt"
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

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

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
			// hash값이 없고, height가 0인 블록체인을 생성
			b = &blockchain{"", 0}

			// search for checkpoint on the db
			checkpoint := db.Checkpoint()

			// if checkpoint exist, decode b from bytes (b는 bytes로 저장되어있음)
			if checkpoint == nil {
				// checkpoint가 없음 즉, db 자체가 없음
				b.AddBlock("Genesis")
			} else {
				// decode b from bytes
				b.restore(checkpoint)
			}
		})
	}

	fmt.Printf("NewestHash: %s, Height:%d\n", b.NewestHash, b.Height)
	return b
}
