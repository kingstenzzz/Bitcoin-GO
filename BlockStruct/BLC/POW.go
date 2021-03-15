package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const DiffucltyBits = 8

type ProofOfWork struct {
	Block          *Block
	ActualTimespan int64
	target         *big.Int
}

func NewProofWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//DiffucltyBits就是前面多少个0
	target = target.Lsh(target, 256-DiffucltyBits)
	return &ProofOfWork{Block: block, target: target}
}

//比较哈希
func (proofOfWork *ProofOfWork) Run() ([]byte, int) {
	var nonce = 0
	var hash []byte
	var hashInt big.Int
	for {
		dataBytes := proofOfWork.prepareData(int64(nonce))
		tmp := sha256.Sum256(dataBytes)
		hash = tmp[:]
		hashInt.SetBytes(hash)
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Printf("block :%d hash count %d\n", proofOfWork.Block.Height, nonce)
	return hash[:], nonce
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heighyByte := IntToHex(pow.Block.Height)
	data = bytes.Join([][]byte{ //HEX添加
		heighyByte,
		timeStampBytes,
		pow.Block.PreBlockHash,
		pow.Block.data,
		IntToHex(nonce),
		IntToHex(DiffucltyBits),
	}, []byte{})
	return data
}
