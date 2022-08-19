package blockchain

import (
	"sync"

	"github.com/baaami/blockcoin/db"
	"github.com/baaami/blockcoin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5 // n개의 block 간 시간을 비교하기 위한 기준 n
	blockInterval      int = 2 // block 1개가 생성되는데 걸리는 목표 시간
	allowedRange       int = 2
)

type blockchain struct {
	// 1. 새로운 블록을 만들기 위해서 마지막 해쉬를 알아야함
	NewestHash string `json:"newestHash"`
	// 2. Height를 알아야함
	Height int `json:"height"`

	CurrentDifficulty int `json:"currentDifficulty"`
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

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

/*
	모든 block을 획득
*/
func (b *blockchain) Blocks() []*Block {
	var blocks []*Block

	// 1. NewestHash로 마지막 block을 찾는다
	hashCursor := b.NewestHash

	// 2. 마지막 block부터 prevHash를 사용하여 첫번째 블럭까지 모두 append
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	// difficultyInterval 개의 블록을 생성하는데 얼마나 걸렸는지 알아야함
	// 너무 오래 걸리면 difficulty를 줄이고 너무 적게 걸리면 늘려야함

	// 최신 block을 찾음
	// 10번일 경우 5번 block부터 5개의 block이 생성되는 동안 걸린 시간을 확인

	// 모든 block 획득
	allBlocks := b.Blocks()
	// 최신 block 획득
	newestBlock := allBlocks[0]

	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval

	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

// mining의 difficulty를 결정하는 함수
func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// recalculate the difficulty
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// 모든 거래 내역 출력 중 address의 출력만 가져와서 슬라이스로 반환하는 함수
func (b *blockchain) UTxOutsByAddress(address string) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)
	for _, block := range b.Blocks() {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Owner == address {
					creatorTxs[input.TxID] = true
				}
			}

			for index, output := range tx.TxOuts {
				if output.Owner == address {
					if _, ok := creatorTxs[tx.ID]; !ok {
						uTxOuts = append(uTxOuts, &UTxOut{tx.ID, index, output.Amount})
					}
				}
			}
		}
	}

	return uTxOuts
}

func (b *blockchain) BalanceByAddress(address string) int {
	txOuts := b.UTxOutsByAddress(address)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

// singleton 인스턴스 획득
func Blockchain() *blockchain {
	if b == nil {
		// 처음 초기화하는 부분은 반드시 1번만 수행되도록
		once.Do(func() {
			// hash값이 없고, height가 0인 블록체인을 생성
			b = &blockchain{
				Height: 0,
			}

			// search for checkpoint on the db
			checkpoint := db.Checkpoint()

			// if checkpoint exist, decode b from bytes (b는 bytes로 저장되어있음)
			if checkpoint == nil {
				// checkpoint가 없음 즉, db 자체가 없음
				b.AddBlock()
			} else {
				// decode b from bytes
				b.restore(checkpoint)
			}
		})
	}
	return b
}
