package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type Block struct {
	PrevBlockHash []byte
	Data          []byte
	TimeStamp     int64
	Hash          []byte
}

func NewBlock(data string, prev []byte) *Block {
	b := &Block{
		PrevBlockHash: prev,
		Data:          []byte(data),
		TimeStamp:     time.Now().Unix(),
	}

	header := bytes.Join([][]byte{b.Data, b.PrevBlockHash, []byte(strconv.FormatInt(b.TimeStamp, 10))}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]

	return b
}

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	c := &BlockChain{}
	c.AppendBlock(GeneisisBlock())
	return c
}

func (b *BlockChain) AppendBlock(block *Block) {
	b.Blocks = append(b.Blocks, block)
}

func GeneisisBlock() *Block {
	return NewBlock("genisis", []byte{})
}

func main() {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-40))
	fmt.Println(target)
}
