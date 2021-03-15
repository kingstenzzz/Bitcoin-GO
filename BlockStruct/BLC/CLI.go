package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BC *BlockChian
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("createblockchain--创建区块链")
	fmt.Println("addblock -data DA--添加区块")
	fmt.Println("printchain --输出区块链信息")

}
func (cli *CLI) createBlockChain() {
	CteateBlockChain()
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock([]byte(data))
}
func (cli *CLI) printChain() {
	cli.BC.PrintChain()
}

func IsValidArgs() {
	if len(os.Args) <= 1 {
		fmt.Printf(os.Args[1])
		fmt.Printf("命令无效")
		PrintUsage()
		os.Exit(1)
	}

}

//命令行运行函数
func (cli *CLI) Run() {
	IsValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ContinueOnError)
	createBlockChianCmd := flag.NewFlagSet("createblockchain", flag.ContinueOnError)

	flagAddBlock := addBlockCmd.String("data", "send 100 btc to xxx", " 区块数据") //参数

	//解析命令
	switch os.Args[1] {

	case "addblock":
		fmt.Println("addblock")
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse addBlockCmd faild %v \n", err)
		}
	case "printchain":
		fmt.Println("printchain")
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printChainCmd faild %v \n", err)

		}
	case "createblockchain":
		fmt.Println("createBlockchain")
		if err := createBlockChianCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse createBlockChianCmd faild %v \n", err)

		}
	default:
		log.Printf("parse  faild  \n")

		//PrintUsage()

		os.Exit(1)
	}
	//解析命令行
	if addBlockCmd.Parsed() {
		blockchain := ReturnBlockOBJ()
		blockchain.AddBlock([]byte(*flagAddBlock))
		//cli.addBlock(*flagAddBlock)
	}
	if printChainCmd.Parsed() {
		blockchain := ReturnBlockOBJ()
		blockchain.PrintChain()
	}
	if createBlockChianCmd.Parsed() {
		cli.createBlockChain()
	}

}
