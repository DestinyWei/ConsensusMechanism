package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

type MinerTime struct {
	minerIndex int
	totalTime  int64
}

var start int64
var end int64

func AddMinerData(minerDatas *[]MinerTime, minerData *MinerTime) {
	*minerDatas = append(*minerDatas, *minerData)
}

func CorrectMiner(Miners *[]Miner, Dif int64, tradeData string) int {
	var minTime int64 = INT64_MAX
	var correctMiner int
	var MinerData []MinerTime
	for i := 0; i < len(*Miners); i++ {
		start = time.Now().UnixNano()
		// 最小持币量为2才能挖矿
		time.Sleep(1)
		if (*Miners)[i].num >= 2 {
			success := Pos((*Miners)[i], Dif, tradeData)
			if success == true {
				end = time.Now().UnixNano()
				MinerDataDemo := MinerTime{
					minerIndex: i,
					totalTime:  end - start,
				}
				AddMinerData(&MinerData, &MinerDataDemo)
			}
		}
	}
	if MinerData != nil {
		fmt.Println(MinerData)
		for j, _ := range MinerData {
			if MinerData[j].totalTime < minTime {
				minTime = MinerData[j].totalTime
				correctMiner = MinerData[j].minerIndex
			}
		}
		(*Miners)[correctMiner].coinAge = 0
		return correctMiner
	}
	return -1

}

// Pos  传入Miners数组，当前难度值Dif和一个string类型变量tradeData，内设一个int变量timeCounter，从0递增到Intmax，
// hash值为SHA256(SHA256(tradeData|timeCounter))，循环内遍历Miners数组，目标值target=Dif乘当前Miner的币龄，
// 要求hash小于target，返回满足要求的第一个Miner的序号并清空这个Miner的币龄，一旦满足要求则退出整个循环
func Pos(Miners Miner, Dif int64, tradeData string) bool {
	var timeCounter int
	var realDif int64
	realDif = int64(MinProbably)
	if realDif+Dif*Miners.coinAge > int64(MaxProbably) {
		realDif = MaxProbably
	} else {
		realDif += Dif * Miners.coinAge
	}

	target := big.NewInt(1)
	// 数据长度为8位
	// 需求：需要满足前两位为0，才能解决问题
	// 1 * 2 << (8-2) = 64
	// 0100 0000
	// 00xx xxxx
	// 32 * 8
	target.Lsh(target, uint(realDif))
	for timeCounter = 0; timeCounter < INT64_MAX; timeCounter++ {
		hash := sha256.Sum256([]byte(tradeData + string(timeCounter)))
		hash = sha256.Sum256(hash[:])
		var hashInt big.Int
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			return true
		}
	}
	return false
}
