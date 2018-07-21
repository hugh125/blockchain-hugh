package main

import (
	"time"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Version int64
	PrevBlockHash []byte
	Hash []byte
	TimeStamp int64
	TargetBits int64
	Nonce int64
	MerKerRoot []byte

	Data []byte
}

/*
@构造函数
 */
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		//Hash:
		TimeStamp:time.Now().Unix(),
		TargetBits:10,
		Nonce:5,
		MerKerRoot:[]byte{},

		Data:[]byte(data),
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash(){
	tmp := [][]byte{
		//实现int类型转为byte类型的工具函数
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		IntToByte(block.TargetBits),
		IntToByte(block.Nonce),
		block.Data}

	//将区块的各个字段连城一个切片，使用[]byte{}进行连接，目的是避免原区块的信息
	data := bytes.Join(tmp, []byte{})
	//将区块进行sha256哈希算法，返回值为[32]byte数组，不是切片
	hash := sha256.Sum256(data)
	block.Hash = hash[:]//有数组转为切片
}

//创建比特币的创世区块，它的前一个区块的hash为空
func NewGenesisBlock() *Block{
	return NewBlock("This is a genesisBlock", []byte{})
}