package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器
type BlockChainIierator struct {
	DB          *bolt.DB
	currentHash []byte
}

//实现迭代函数，获取区块
func (blockchain *BlockChian) Iterator() *BlockChainIierator {
	return &BlockChainIierator{DB: blockchain.DB, currentHash: blockchain.Tip}
}

func (bcit *BlockChainIierator) Next() *Block {
	var block *Block
	err := bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(bcit.currentHash)
			block = Deserialize(currentBlockBytes)
			bcit.currentHash = block.PreBlockHash
		}
		return nil
	})
	if err != nil {
		log.Printf("iterator the db failed %v\n", err)
	}
	return block
}
