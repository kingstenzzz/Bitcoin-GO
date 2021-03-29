package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
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
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]

}

//生成普通转账
func NewSimpleTransaction(from string, to string, amount int, bc *BlockChian) *Transaction {
	var txInputs []*TxInput
	var txOutputs []*TxOutput
	//blockchian := ReturnBlockOBJ()
	money, spendableUTXOic := bc.FindSpendableUTXO(from, amount)
	fmt.Printf("momey %v", money)
	//输入
	for txHash, indexArray := range spendableUTXOic {
		txHashBytes, err := hex.DecodeString(txHash)
		if err != nil {
			fmt.Printf("deconde spendable string to byte faild%v", err)
		}
		for _, index := range indexArray {
			txInput := &TxInput{TxHash: txHashBytes, Vout: index, ScriptSig: from}
			txInputs = append(txInputs, txInput)
		}

	}

	txOutput := &TxOutput{Value: amount, ScriptPubkey: to}
	txOutputs = append(txOutputs, txOutput)
	//找零
	if money > amount {
		txOutput = &TxOutput{money - amount, from}
		txOutputs = append(txOutputs, txOutput)

	} else {
		fmt.Printf("余额不足..")

	}

	tx := Transaction{Vins: txInputs, Vouts: txOutputs, TxHash: nil}
	tx.HashTransaction()
	return &tx

}

//判断是否coinbase交易
func (tx *Transaction) isCoinbaseTransaction() bool {
	return tx.Vins[0].Vout == -1 && len(tx.Vins[0].TxHash) == 0

}
