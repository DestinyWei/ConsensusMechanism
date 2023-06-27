package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Miner 创建一种名为Miner的结构体，包含miner的地址和持币数量，以及记录的币龄
type Miner struct {
	addr    []byte
	num     int64
	coinAge int64
}

// createMiner 初始化Miner的函数，默认addr为用sha256方法对字符串miner和现在的时间拼接后的字符串处理后的结果,num为0，coinAge为0
func createMiner() *Miner {
	temp := sha256.Sum256([]byte("miner" + time.Now().String()))
	miner := Miner{
		addr:    temp[:],
		num:     0,
		coinAge: 0,
	}
	return &miner
}

// InitMiners 初始化Miners数组的函数，调用AddMiner函数，生成一个Miner，然后将其添加到Miners数组中
func InitMiners() []Miner {
	miner := createMiner()
	Miners := []Miner{*miner}
	return Miners
}

// AddMiner 传入一个Miner和Miners数组，将miner添加到Miners数组中
func AddMiner(miner Miner, Miners *[]Miner) {
	*Miners = append(*Miners, miner)
}

// UpdateMiners 更新Miners数组函数，传入Coins数组和Miners数组，遍历Coins数组，将Coins数组中的币的矿工序号与Miners数组中的矿工序号相同的矿工的币龄加上（现在的时间-Coin的时间戳）*Coin的数量
func UpdateMiners(Coins *[]Coin, Miners *[]Miner) []Miner {
	for i := 0; i < len(*Coins); i++ {
		index := (*Coins)[i].MinerIndex
		(*Miners)[index].coinAge += (time.Now().Unix() - (*Coins)[i].Time) * (*Coins)[i].Num
		(*Coins)[i].Time = time.Now().Unix()
	}
	return *Miners
}

func AddMiners() {
	var MinerNum int
	fmt.Print("请输入创建矿工的数量：")
	fmt.Scanf("%d", &MinerNum)
	for i := 0; i < MinerNum; i++ {
		AddMiner(*createMiner(), &Miners)
	}
}
