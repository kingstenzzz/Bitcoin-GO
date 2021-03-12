package main

import (
	"../BlockStruct/BLC"
	"fmt"
)

func main() {
	//block := BLC.NewBlock(1,nil,[]byte("First Block"))
	bc := BLC.CteateBlockChain()

	bc.AddBlock(bc.Blocks[0].Height+1, bc.Blocks[0].Hash, []byte("secend "))
	bc.AddBlock(bc.Blocks[1].Height+1, bc.Blocks[1].Hash, []byte("three "))

	for _, block := range bc.Blocks {
		fmt.Printf("Hash: %X\n", block.Hash)

	}

}
