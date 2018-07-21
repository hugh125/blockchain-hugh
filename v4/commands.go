package main

import "fmt"

func (cli *CLI) AddBlock(data string){
	cli.bc.AddBlock(data)
	println("AddBlock Successed!")
}

func (cli *CLI) PrintChain() {
	bc := cli.bc
	it := bc.Iterator()
	for{
		//取回当前hash指向的block， 将hash值前移
		block := it.Next()
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Version: %x\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("TargetBits: %x\n", block.TargetBits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("MerKerRoot: %x\n", block.MerKerRoot)


		if len(block.PrevBlockHash) == 0{
			break
		}
	}

}
