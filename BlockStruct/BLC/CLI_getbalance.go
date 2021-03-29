package BLC

import (
	"fmt"
	"os"
)

//查询指定地址得UTXO
func (cli *CLI) getBalance(from string) {
	if !dbExist() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := ReturnBlockOBJ()
	blockchain.UnUTXOS(from)
	defer blockchain.DB.Close()
	amount := blockchain.getBalance(from)
	fmt.Printf("\t地址[%s]的余额：[%d]\n", from, amount)

}
