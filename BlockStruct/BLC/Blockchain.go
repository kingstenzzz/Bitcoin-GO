package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

//数据库名字
const dbName = "block.db"
const blockTableName = "blocks"

type BlockChian struct {
	Blocks []*Block
	DB     *bolt.DB //数据库对象
	Tip    []byte
}

//添加区块
func (bc *BlockChian) AddBlock(data []byte) {
	bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取最新取快递饿hash
			blockByte := b.Get(bc.Tip)
			latest_block := Deserialize(blockByte)
			//数据库出来的数据需要反序列化
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, data)
			bc.Blocks = append(bc.Blocks, newBlock) //添加到结构体
			//存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Printf("insert new block to db faild %v", err)
			}
			//存取最新的哈希
			err = b.Put([]byte("1"), newBlock.Hash)
			if err != nil {
				log.Printf("update the latest block to db faild %v", err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil

	})
}

//创建区块链
func CteateBlockChain() *BlockChian {
	var latestNlockHash []byte
	var fiestBlock *Block
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panic("create db[%s] faild %v\n", dbName, err)
	}

	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//空
			b, err := tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Printf("create bucket [%s]faild %v \n", blockTableName, err)
			}
			fiestBlock = CreateGenesisBlock([]byte("GenesisBlock"))
			//先发数据序列化后才能够存去数据库
			err = b.Put(fiestBlock.Hash, fiestBlock.Serialize())
			if err != nil {
				log.Panicf("insert the genensis block faild %v \n", err)
			}
			latestNlockHash = fiestBlock.Hash
			err = b.Put([]byte("1"), fiestBlock.Hash)
			if err != nil {
				log.Printf("save the hash of genesis block faild %v \n", err)
			}
		}
		return nil

	})
	if nil != err {
		log.Printf("update db faild %v \n", err)
	}
	return &BlockChian{DB: db, Tip: latestNlockHash, Blocks: []*Block{fiestBlock}}
}

//便利区块链
func ReturnTheChain(bc *BlockChian) {
	bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取最新取快递饿hash
			blockByte := b.Get(bc.Tip)
			for {
				if blockByte == nil {
					break
				}
				last_block := Deserialize(blockByte)
				fmt.Println("")
				//bc.Blocks = append(bc.Blocks, last_block) //添加到结构体
				blockByte = b.Get(last_block.PreBlockHash)
			}
			//数据库出来的数据需要反序列化
		}
		return nil
	})

}

func (bc *BlockChian) PrintChain() {
	fmt.Println("完整区块信息...")
	var curBlock *Block
	bcit := bc.Iterator() //获取迭代对象
	//var  currentHash = bc.Tip
	for {
		fmt.Println("-----")
		curBlock = bcit.Next()
		fmt.Printf("\tHash %x \n", curBlock.Hash)
		fmt.Printf("\tPreBlockHash %x \n", curBlock.PreBlockHash)
		fmt.Printf("\tdata %x \n", curBlock.data)
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}
