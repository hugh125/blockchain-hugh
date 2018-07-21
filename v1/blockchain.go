package main

import "os"

//改造区块结构，使用数组来储存所有的区块
type BlockChain struct {
	blocks []*Block
}

//创建区块链实例，并且添加第一个创世快
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//添加区块操作
func (bc *BlockChain)AddBlock(data string){
	//校验数组元素个数，避免出现访问越界情况！！！
	if len(bc.blocks) <= 0{
		os.Exit(1)
	}

	//取出最后一个区块，目的是得到其hash值
	lastBlock := bc.blocks[len(bc.blocks) - 1]
	prevBlockHash := lastBlock.Hash
	//改造一个新的将要添加到区块链的区块
	newBlock := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, newBlock)
}
