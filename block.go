package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block header
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns a block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns a genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
