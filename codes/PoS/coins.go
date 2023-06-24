package main

import (
	"crypto/rand"
	"math/big"
	"time"
)

// 创建一种名为Coin的结构体，包含币的数量，矿工序号，币的时间戳

type Coin struct {
	Num        int64
	MinerIndex int64
	Time       int64
}

// 生成新coin函数，传入矿工序号，返回一个coin

func NewCoin(MinerIndex int64, Miners []Miner) Coin {
	n, _ := rand.Int(rand.Reader, big.NewInt(4))
	coin := Coin{
		Num:        1 + n.Int64(),
		MinerIndex: MinerIndex,
		Time:       time.Now().Unix(),
	}
	Miners[MinerIndex].num += coin.Num
	return coin
}

// 初始化coins数组函数，调用NewCoin函数，生成一个coin，然后将其添加到coins数组中
func InitCoins(Miners []Miner) []Coin {
	coin := NewCoin(0, Miners)
	Coins := []Coin{coin}
	return Coins
}

// 传入新coin和coins数组，将其添加到coins数组中并保存，无返回值
func AddCoin(coin Coin, Coins *[]Coin) {
	*Coins = append(*Coins, coin)
}
