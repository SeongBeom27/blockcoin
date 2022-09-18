package db

import (
	"github.com/baaami/blockcoin/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		// init db
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		db = dbPointer
		utils.HandleErr(err)

		err = db.Update(func(tx *bolt.Tx) error {
			// check two bucket exist
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func SaveBlock(hash string, data []byte) {
	// fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

/*
	@brief	blockchain에 checkpoint가 있는지 없는지를 알 수 있음
*/
func Checkpoint() []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
