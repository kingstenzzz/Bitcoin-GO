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

	for _, block := range bc.Blocks {
		fmt.Printf("Hash: %X\n", block.Hash)

	}

}
