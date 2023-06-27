package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

var (
	// Nonce循环上限
	maxNonce = math.MaxInt64
)

// 难度值
const targetBits = 24

// ProofOfWork pow结构
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork 创建pow结构
func NewProofOfWork(b *Block) *ProofOfWork {
	// target为最终难度值
	target := big.NewInt(1)
	// target为1向左位移256-24（挖矿难度）
	target.Lsh(target, uint(256-targetBits))
	// 生成pow结构
	pow := &ProofOfWork{b, target}
	return pow
}

// Run 挖矿运行
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing %s,maxNonce=%d\n", pow.block.Data, maxNonce)
	for nonce < maxNonce {
		// 数据准备
		data := pow.prepareData(int64(nonce))
		// 计算哈希
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		// 按字节比较，hashInt cmp小于0代表找到目标Nonce
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")

	return nonce, hash[:]
}

// prepareData 准备数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int2Hex(pow.block.Timestamp),
			Int2Hex(int64(targetBits)),
			Int2Hex(nonce),
		},
		[]byte{},
	)

	return data
}

// Int2Hex 将int64写入[]byte
func Int2Hex(num int64) []byte {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, num)
	return buff.Bytes()
}

// Validate 校验区块正确性
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
