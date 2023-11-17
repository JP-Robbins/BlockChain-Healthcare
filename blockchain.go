package main

// create the method that adds a new block to the blockchain
func (blockchain *blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	NewBlock := NewBlock(data, PreviousBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, NewBlock)
}

func NewBlockchain() *blockchain {
	return &blockchain{[]*Block{NewGenesisBlock()}}
}