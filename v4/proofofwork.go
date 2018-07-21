package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 1

type ProofOfWork struct {
	block     *Block
	targetBit *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	var IntTarget = big.NewInt(1)
	IntTarget.Lsh(IntTarget, uint(256-targetBits))
	return &ProofOfWork{block: block, targetBit: IntTarget}
}

func (pow *ProofOfWork) PrepareRawData(nonce int64) []byte {
	block := pow.block
	tmp := [][]byte{
		//实现int类型转为byte类型的工具函数
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKerRoot,
		IntToByte(nonce),
		IntToByte(targetBits),
		block.Data}

	//将区块的各个字段连城一个切片，使用[]byte{}进行连接，目的是避免原区块的信息
	data := bytes.Join(tmp, []byte{})

	return data
}

func (pow *ProofOfWork) run() (int64, []byte) {
	var nonce int64
	var hash [32]byte
	var hashInt big.Int

	println("Begin Mining...")
	fmt.Printf("target hash   : %X\n", pow.targetBit.Bytes())
	for nonce < math.MaxInt64 {
		//将所有字段连接为字节数组，便于生成256字节
		data := pow.PrepareRawData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.targetBit) == -1 {
			fmt.Printf("Fount Hash: %x\n", hash)
			break
		} else {
			//fmt.Printf("current Hash: %x\n", hash)
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow ProofOfWork)IsValid()bool{
	data := pow.PrepareRawData(pow.block.Nonce)
	hash :=sha256.Sum256(data)
	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.targetBit) == -1
}
