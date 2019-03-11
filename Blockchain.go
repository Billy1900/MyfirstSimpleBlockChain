package BLC

type Blockchain struct {
	Blocks []*Block//存储有序区块
}

func CreateBlockChainWithGenesisBlock(data string) *Blockchain {
	genesisBlock := CreateGenesisBlock(data)
	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlockToBlockChain(data string,height int64,prevhash []byte){
	newBlock := NewBlock(data,prevhash,height)
	bc.Blocks = append(bc.Blocks,newBlock)//添加到切片中
}