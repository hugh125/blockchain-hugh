package main

import "fmt"

func main(){
	bc := NewBlockChain()
	bc.AddBlock("000")
	bc.AddBlock("001")
	//bc.AddBlock("002")
	//bc.AddBlock("003")

	for i, block := range(bc.blocks){
		println("==========block num: ", i)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Version: %x\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("TargetBits: %x\n", block.TargetBits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("MerKerRoot: %x\n", block.MerKerRoot)
	}
}
