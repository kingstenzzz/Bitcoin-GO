package main

import "./BLC"

func main() {
	//block := BLC.NewBlock(1,nil,[]byte("First Block"))
	//bc  := BLC.CteateBlockChain()
	//bc := BLC.ReturnBlockOBJ()
	//bc.AddBlock([]byte("secend"))
	//bc.AddBlock([]byte("three"))
	//bc.PrintChain()

	//for _, block := range bc.Blocks {
	//fmt.Printf("Block :%d Hash: %X\n", block.Height, block.Hash)
	//}

	cli := BLC.CLI{}
	cli.Run()

}
