package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC To Ivan")
	bc.AddBlock("Send 2 More BTC to Ivan")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println()
	}
}
