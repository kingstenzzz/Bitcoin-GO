package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Transaction struct {
	TxHash []byte      ///交易哈希
	Vins   []*TxInput  //输入列表
	Vouts  []*TxOutput //输出列表
}

func NewCoinbaseTransaction(address string) *Transaction {
	var txConbase *Transaction
	txinput := &TxInput{vout: -1, ScriptSig: "system awward", TxHash: []byte{}}
	txoutPut := &TxOutput{value: 10, ScriptPubkey: address}
	txConbase = &Transaction{TxHash: nil, Vins: []*TxInput{txinput}, Vouts: []*TxOutput{txoutPut}}
	txConbase.HashTransaction()
	return txConbase

}

//生成交易哈希
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(tx); err != nil {
		log.Printf("tx HASH encode failed %v \n", err)
	}

}
