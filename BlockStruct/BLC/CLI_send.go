package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) sendCoinTo(from, to, amount []string) {
	if !dbExist() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockhcain := ReturnBlockOBJ()
	defer blockhcain.DB.Close()
	blockhcain.MineNewBlock(from, to, amount)

}
