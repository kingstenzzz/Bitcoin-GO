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
	fmt.Println("-from From -to TO -amount AMOUNT--发起转账")
	//查询余额
	fmt.Println("getbalance -address From --查询指定地址余额")

}

func (cli *CLI) addBlock(txs []*Transaction) {
	if !dbExist() {
		fmt.Println("数据库不存在")
		os.Exit(1)

	}
	cli.BC.AddBlock(txs)
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
	createBlockChianCmd := flag.NewFlagSet("createblockchain -address address", flag.ContinueOnError)
	sendCoinCmd := flag.NewFlagSet("send -from address -to  adress", flag.ContinueOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance -from address", flag.ContinueOnError)
	///
	flagAddBlock := addBlockCmd.String("data", "send 100 btc to xxx", " 区块数据")         //参数
	flagCreateBlockchain := createBlockChianCmd.String("address", "kingsten", " 矿工地址") //参数
	//发起交易
	flagSendFromArg := sendCoinCmd.String("from", "", " 源地址")    //参数
	flagSendToArg := sendCoinCmd.String("to", "", " 接收地址")       //参数
	flagSendAmountArg := sendCoinCmd.String("amount", "", " 数量") //参数
	//查询余额

	flagGetBalanceArg := getBalanceCmd.String("from", "", " 源地址") //参数

	fmt.Println(os.Args)

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
	case "send":
		fmt.Println("send coin")
		if err := sendCoinCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("send coin faild %v \n", err)
		}
	case "getbalance":
		fmt.Println("get Balace")
		if err := getBalanceCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("get balance faild %v \n", err)
		}
	default:
		log.Printf("parse  faild  \n")

		//PrintUsage()

		os.Exit(1)
	}
	//解析命令行
	if addBlockCmd.Parsed() {
		blockchain := ReturnBlockOBJ()
		fmt.Printf(*flagAddBlock)
		blockchain.AddBlock([]*Transaction{}) /////
	}
	if printChainCmd.Parsed() {
		blockchain := ReturnBlockOBJ()
		blockchain.PrintChain()
	}
	if createBlockChianCmd.Parsed() {
		if *flagCreateBlockchain == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockChain(*flagCreateBlockchain)
	}
	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceArg)
	}

	if sendCoinCmd.Parsed() {
		if *flagSendFromArg == "" {
			fmt.Println("源地址不能空")
			os.Exit(1)
		}
		if *flagSendToArg == "" {
			fmt.Println("接收不能空")
			os.Exit(1)
		}
		if *flagSendAmountArg == "" {
			fmt.Println("数量")
			os.Exit(1)
		}

		fmt.Printf("From %s", JSONToSlice(*flagSendFromArg))
		fmt.Printf("to %s", JSONToSlice(*flagSendToArg))
		fmt.Printf("amount %s", JSONToSlice(*flagSendAmountArg))
		cli.sendCoinTo(JSONToSlice(*flagSendFromArg), JSONToSlice(*flagSendToArg), JSONToSlice(*flagSendAmountArg))
	}

}
