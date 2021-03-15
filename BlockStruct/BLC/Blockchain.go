package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
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
func (blc *BlockChian) AddBlock(data []byte) {
	blc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取最新取快递饿hash
			blockByte := b.Get(blc.Tip)
			latest_block := Deserialize(blockByte)
			//数据库出来的数据需要反序列化
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, data)
			//fmt.Printf("写入数据 %s",data)
			blc.Blocks = append(blc.Blocks, newBlock) //添加到结构体
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
			blc.Tip = newBlock.Hash
		}
		return nil

	})
}

//err属于IsExist错误
//判断数据库是否存在
func dbExist() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true

}

//创建区块链

func CteateBlockChain() *BlockChian {
	//先检测区块链是否已经存在
	if dbExist() {
		fmt.Printf("区块链已存在\n")
		os.Exit(1)
	}
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

//遍历区块链
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
				fmt.Printf("Height%v \n", last_block.Height)
				fmt.Printf("%v\n", last_block.data)
				//bc.Blocks = append(bc.Blocks, last_block) //添加到结构体
				blockByte = b.Get(last_block.PreBlockHash)
			}
			//数据库出来的数据需要反序列化
		}
		return nil
	})

}

func (blc *BlockChian) PrintChain() {
	fmt.Println("完整区块信息...")
	var curBlock *Block
	bcit := blc.Iterator() //获取迭代对象
	//var  currentHash = bc.Tip
	for {
		fmt.Println("-----")
		curBlock = bcit.Next()
		fmt.Printf("\tHeight %x \n", curBlock.Height)
		fmt.Printf("\tHash %x \n", curBlock.Hash)
		fmt.Printf("\tPreBlockHash %x \n", curBlock.PreBlockHash)
		fmt.Printf("Data: %v \n", curBlock.data)
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

//返回一个Blockchain 对象
func ReturnBlockOBJ() *BlockChian {
	//DB
	//TIP
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panic("open db[%s] faild %v\n", dbName, err)
	}
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			tip = b.Get([]byte("1"))
		}
		return nil

	})
	if nil != err {
		log.Println("查找区块链失败")
	}
	return &BlockChian{DB: db, Tip: tip}

}
