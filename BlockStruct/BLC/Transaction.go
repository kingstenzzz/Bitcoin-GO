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

//创世区块输入为--1
func NewCoinbaseTransaction(address string) *Transaction {
	var txConbase *Transaction
	txinput := &TxInput{Vout: -1, ScriptSig: "block award", TxHash: []byte{}}
	txoutPut := &TxOutput{Value: 10, ScriptPubkey: address}
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

//生成普通转账
func NewSimpleTransaction(from string, to string, ammount int) *Transaction {
	var txInputs []*TxInput
	var txOutputs []*TxOutput
	//blockchian := ReturnBlockOBJ()

	txInput := &TxInput{}

	txInputs = append(txInputs, txInput)
	txOutput := &TxOutput{}
	txOutputs = append(txOutputs, txOutput)
	tx := Transaction{Vins: txInputs, Vouts: txOutputs, TxHash: nil}
	tx.HashTransaction()
	return &tx

}
