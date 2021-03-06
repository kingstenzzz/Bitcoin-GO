package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//区块结构
type Block struct {
	TimeStamp    int64
	PreBlockHash []byte
	Hash         []byte
	Height       int64
	Txs          []*Transaction //交易数据
	Nonce        int64          //pow生成的哈希变化值
}

func NewBlock(height int64, preBlockHash []byte, txs []*Transaction) *Block {
	var block Block
	block.TimeStamp = time.Now().Unix()
	block.PreBlockHash = preBlockHash
	block.Height = height
	block.Txs = txs
	//HASH是根据当前的参数生成的
	block.SetHash()
	pow := NewProofWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

//**生成HASH*/
func (b *Block) SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heighyByte := IntToHex(b.Height)
	blockByte := bytes.Join([][]byte{ //HEX添加
		heighyByte,
		timeStampBytes,
		b.PreBlockHash,
		b.HashTransaction(),
	}, []byte{})

	hash := sha256.Sum256(blockByte)
	b.Hash = hash[:]
	//(b.Hash) = sha256.Sum256(nil)
}

func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(1, nil, txs)

}

//区块结构序列化
func (block *Block) Serialize() []byte {
	var bufer bytes.Buffer
	encoder := gob.NewEncoder(&bufer)
	//编码序列化
	if err := encoder.Encode(block); nil != err {
		log.Printf("serialize the blocl to []byte faild %v \n", err)
	}
	return bufer.Bytes()
}

func Deserialize(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); nil != err {
		log.Printf("deserialize byte to block fail......%v\n", err)
	}

	return &block

}

func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash) //把交易添加到UTXO集合

	}
	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]

}
