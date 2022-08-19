package blockchain

import (
	"time"

	"github.com/baaami/blockcoin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx // confirm 받기 전의 transcation들
}

// 메모리에서만 사용하여 전역으로 선언 및 초기화
var Mempool *mempool = &mempool{}

// Transaction
type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txins"`
	TxOuts    []*TxOut `json:"txouts"`
}

// 거래 내역(transactions) 들에 대한 hash 값으로 id를 생성
func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	TxID  string `json:"txId"` // input 으로 사용하고 있는 output을 생성한 transaction을 의미
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type UTxOut struct {
	TxID   string `json:"txId"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {

}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("baami", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("baami")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
