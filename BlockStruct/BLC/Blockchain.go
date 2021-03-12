package BLC

type BlockChian struct {
	Blocks []*Block
}

//添加区块
func (bc *BlockChian) AddBlock(height int64, preBlockHash []byte, data []byte) {
	var newBlock *Block
	newBlock = NewBlock(height, preBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//创建区块链
func CteateBlockChain() *BlockChian {
	firstblock := CreateGenesisBlock([]byte("GenesisBlock"))
	blockChain := BlockChian{[]*Block{firstblock}}
	return &blockChain

}
