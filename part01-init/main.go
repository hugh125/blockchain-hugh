package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to block1")
	bc.AddBlock("Send 1 BTC to block2")

	for _, block := range bc.blocks {
		fmt.Printf("PecvHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		println()
	}
}
