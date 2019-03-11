package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//block
type Block struct {
	Height int64
	PrevBlockHash []byte
	Data []byte
	TimeStamp int64
	Hash []byte
}

//new block hash
func NewBlock(data string, PrecBlockHash []byte,height int64) *Block {
	block := &Block{height, PrecBlockHash,[]byte(data),time.Now().Unix(),nil}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	heightByte := IntToHex(block.Height)//height
	timeString := strconv.FormatInt(block.TimeStamp,2)//timestamp->2进制
	timebytes := []byte(timeString)
	//catch
	blockBytes := bytes.Join([][]byte{
		heightByte,
		block.PrevBlockHash,
		block.Data,
		timebytes},[]byte{}) //[]byte{}为合并后的类型
	//->hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]//切片操作引用整个数组
}

//genesis block
func CreateGenesisBlock (data string) *Block {
	return NewBlock(data,make([]byte,32,32),0)//make用法：声明32byte大小，预留32大小空间
}