package main

import (
	"fmt"
)

func main() {
	newblockchain := NewBlockchain()

	newblockchain.AddBlock("patient one")
	newblockchain.AddBlock("patient two")

	for i, block := range newblockchain.Blocks {
		fmt.Println("---------------------------------------------")
		fmt.Printf("Block ID: %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %s\n", block.AllData)
		fmt.Println("---------------------------------------------")

	}
}
