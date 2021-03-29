package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) printChain() {
	if !dbExist() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	cli.BC.PrintChain()

}
