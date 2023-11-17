package main

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	MyBlockHash       []byte
	AllData           []byte
}

type blockchain struct {
	Blocks []*Block
}
