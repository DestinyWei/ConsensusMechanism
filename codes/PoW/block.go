package main

import (
	"time"
)

// Block 定义区块结构
type Block struct {
	Timestamp     int64  // 时间戳
	Data          []byte // 数据域
	PrevBlockHash []byte // 前一块hash
	Hash          []byte // 当前块哈希
	Nonce         int64  // 随机值
}

// NewBlock 创建Block返回Block指针
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 构造block
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	// 挖矿
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	// 设置hash和nonce
	block.Hash = hash
	block.Nonce = int64(nonce)
	return block
}

// NewGenesisBlock 创世区块创建 返回创世块Block指针
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
