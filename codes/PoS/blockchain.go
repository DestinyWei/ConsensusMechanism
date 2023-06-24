package main

import (
	"crypto/sha256"
	"math"
	"time"
)

const (
	dif         = 2
	INT64_MAX   = math.MaxInt64
	MaxProbably = 255
	MinProbably = 235
	MaxCoinAge  = 10
	Minute      = 60
)

// 创建一种名为Block的结构体,包含区块哈希，前区块哈希，区块号，难度值，矿工地址，奖励币数，时间戳
type Block struct {
	Hash      []byte
	PrevHash  []byte
	Height    int64
	Dif       int64
	MinerAddr string
	Reward    Coin
	Timestamp int64
	tradeData string
}

// 初始化函数，生成创世区块，并添加到区块链中

func InitBlockChain(Miners []Miner, Coins []Coin) []Block {
	var bc []Block
	bc = append(bc, GenesisBlock(Miners, Coins))
	return bc
}

// 生成创世区块，默认难度值为1，矿工地址为矿工数组0

func GenesisBlock(Miners []Miner, Coins []Coin) Block {
	temp := sha256.Sum256([]byte("Genesis Block"))
	genesisBlock := Block{
		Hash:      temp[:],
		tradeData: "Genesis Block",
		PrevHash:  []byte(""),
		Height:    1,
		Dif:       0,
		MinerAddr: string(Miners[0].addr),
		Reward:    Coins[0],
		Timestamp: time.Now().Unix(),
	}
	return genesisBlock
}

// 生成区块函数，传入参数为矿工序号，矿工数组，Coin,tradeData,区块数组,新区块的Hash是tradeData的sha256的运算结果，PrevHash是上一个区块的哈希，区块号是上一个区块的区块号加1，难度值是上一个区块的难度值，矿工地址是矿工数组中对应序号的地址，奖励币数是Coin，时间戳是当前时间戳，将新生成的区块添加到区块数组中
func GenerateBlock(MinerNum int, Miners []Miner, coin Coin, tradeData string, bc *[]Block) {
	var newBlock Block
	temp := sha256.Sum256([]byte(tradeData))
	newBlock.Hash = temp[:]
	newBlock.PrevHash = (*bc)[len(*bc)-1].Hash
	newBlock.Height = (*bc)[len(*bc)-1].Height + 1
	newBlock.Dif = (*bc)[len(*bc)-1].Dif
	newBlock.MinerAddr = string(Miners[MinerNum].addr)
	newBlock.Reward = coin
	newBlock.Timestamp = time.Now().Unix()
	newBlock.tradeData = tradeData
	*bc = append(*bc, newBlock)
}
