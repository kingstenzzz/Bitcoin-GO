package main

import (
	"../BlockStruct/BLC"
	"fmt"
)

func main() {
	//block := BLC.NewBlock(1,nil,[]byte("First Block"))
	bc := BLC.CteateBlockChain()

	bc.AddBlock([]byte("secend "))
	bc.AddBlock([]byte("three "))
	//BLC.ReturnTheChain(bc)
	for _, block := range bc.Blocks {
		fmt.Printf("Block :%d Hash: %X\n", block.Height, block.Hash)
	}
}
