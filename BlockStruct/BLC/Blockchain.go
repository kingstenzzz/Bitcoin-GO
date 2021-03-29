package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
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
func (blockchain *BlockChian) AddBlock(txs []*Transaction) {

	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取最新取快递饿hash
			blockByte := b.Get(blockchain.Tip)
			//数据库出来的数据需要反序列化
			latest_block := Deserialize(blockByte)
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, txs)
			//fmt.Printf("写入数据 %s",data)
			blockchain.Blocks = append(blockchain.Blocks, newBlock) //添加到结构体
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
			blockchain.Tip = newBlock.Hash
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

func CteateBlockChain(address string) *BlockChian {
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
			txCoinBase := NewCoinbaseTransaction(address)

			fiestBlock = CreateGenesisBlock([]*Transaction{txCoinBase})

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
				fmt.Printf("%v\n", last_block.Txs)
				//bc.Blocks = append(bc.Blocks, last_block) //添加到结构体
				blockByte = b.Get(last_block.PreBlockHash)
			}
			//数据库出来的数据需要反序列化
		}
		return nil
	})
}

func (blockchain *BlockChian) PrintChain() {
	fmt.Println("完整区块信息...")
	var curBlock *Block
	bcit := blockchain.Iterator() //获取迭代对象
	//var  currentHash = bc.Tip
	for {
		fmt.Println("-----")
		curBlock = bcit.Next()
		fmt.Printf("\tHeight: %d \n", curBlock.Height)
		fmt.Printf("\tHash: %x \n", curBlock.Hash)
		fmt.Printf("\tPreBlockHash :%x \n", curBlock.PreBlockHash)
		for _, tx := range curBlock.Txs {
			fmt.Printf("\ttxHash: %x \n", tx.TxHash)
			for _, vin := range tx.Vins {
				fmt.Printf("\tvin- Hash: %x \n", vin.TxHash)
				fmt.Printf("\tvin -OUT: %d \n", vin.Vout)
				fmt.Printf("\tvin -Sig: %s \n", vin.ScriptSig)
			}
			for _, vout := range tx.Vouts {
				fmt.Printf("\tout- value: %d \n", vout.Value)
				fmt.Printf("\tout - ScriptPubkey:%s \n", vout.ScriptPubkey)

			}

		}
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

//挖矿功能
func (blockchain *BlockChian) MineNewBlock(from, to, amount []string) {
	var txs []*Transaction
	var block *Block
	//生成交易
	for index, address := range from {
		value, _ := strconv.Atoi(amount[index])
		tx := NewSimpleTransaction(address, to[index], value, blockchain)
		txs = append(txs, tx)

	}
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte("1"))
			blockByte := b.Get(hash)
			block = Deserialize(blockByte)

		}
		return nil
	})
	//生成最新区块
	block = NewBlock(block.Height+1, block.Hash, txs)
	//新区块添加到数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			err := b.Put(block.Hash, block.Serialize())
			if nil != err {
				log.Printf("update the new block to DB failed %v\n", err)
			}
			err = b.Put([]byte("1"), block.Hash)
			if err != nil {
				log.Printf("update the latest hash to the DB failed %v \n", err)
			}
			blockchain.Tip = block.Hash

		}
		return nil

	})

}

/*
遍历查找区块链数据库中的每一个区块中的每一个
1属于传入地址
2属于未被花费交易
 1.遍历数据库，把所有得output存入缓存
 2.再次遍历区块链数据库，检查每一个vout是否包含前面得已花费缓存中
*/
//返回所有的输出
func (blockchian *BlockChian) UnUTXOS(address string) []*UTXO {
	//遍历数据库
	bcit := blockchian.Iterator()
	//获取所有已花费输出
	var unUTXOS []*UTXO
	spentTxOutputs := blockchian.SpentOutputs(address)
	for {
		block := bcit.Next()
		//获取每个区块得交易
		for _, tx := range block.Txs {

		work:
			for index, vout := range tx.Vouts {
				//交易索引位置
				//vout当前输出

				if vout.CheckPubkeyWithAdress(address) { //接收是这个地址的输出
					if len(spentTxOutputs) != 0 {
						var isSpentOutput bool
						for txHash, indexArray := range spentTxOutputs {
							//txHash :当前输出所引用得交易哈希
							//indexArray :哈希关联得vout索引列表
							for _, i := range indexArray {
								//这个
								if txHash == hex.EncodeToString(tx.TxHash) && index == i {
									//说明当前交易tx至少已经有输出被其他交易得输入引用
									//index==i 正好是当先输出被其他交易引用
									isSpentOutput = true
									continue work //跳转到最外层循环
									//index ==i 说明正好是当前得输出被其他交易引用
								}

							}

						}
						/*
							type UTXO struct {
								//UTXO
								TxHash	[]byte
								//UTXO在其所属交易得输出列表中的索引
								Index	int
								//Output
								Output	*TxOutput
							}

						*/
						//没有被花
						if !isSpentOutput == false {
							utxo := &UTXO{TxHash: tx.TxHash, Index: index, Output: vout}
							unUTXOS = append(unUTXOS, utxo)
						}

					} else {
						utxo := &UTXO{TxHash: tx.TxHash, Index: index, Output: vout}

						unUTXOS = append(unUTXOS, utxo)

					}

				}

			}
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return unUTXOS

}

//	//获取所有已花费输出
func (blockchian *BlockChian) SpentOutputs(address string) map[string][]int {
	//已花费输出缓存
	spentTXOutputs := make(map[string][]int)
	bcit := blockchian.Iterator()
	for {
		block := bcit.Next()
		for _, tx := range block.Txs {
			//排出coinbase交易
			if !tx.isCoinbaseTransaction() {
				for _, in := range tx.Vins {
					if in.CheckPubkeyWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}

				}
			}

		}
		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}

	}

	return spentTXOutputs

}

func (blockchain *BlockChian) getBalance(address string) int {
	var amount int
	utxos := blockchain.UnUTXOS(address)
	for _, utxo := range utxos {
		amount += utxo.Output.Value
	}
	return amount
}

//查找指定地址得可用UTXO
func (blockchain *BlockChian) FindSpendableUTXO(from string, amount int) (int, map[string][]int) {
	spendableUTXO := make(map[string][]int)
	var value int
	utxos := blockchain.UnUTXOS(from)
	//遍历UTXO
	for _, utxo := range utxos {
		value += utxo.Output.Value
		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash], utxo.Index)

		if value >= amount {
			break
		}
	}
	//所有的遍历完成
	if value < amount {
		fmt.Printf("地址[%s]余额不足，当前余额[d]，转账[余额]", from, value, amount)
	}
	return value, spendableUTXO

}
