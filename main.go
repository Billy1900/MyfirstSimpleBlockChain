package main

import (
	"Mypublicchain/main/BLC"
	"fmt"
)

func main() {
	blockchain := BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	blockchain.AddBlockToBlockChain("Send 100 to zhang",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("Send 200 to luo",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("Send 300 to wang",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	fmt.Println(blockchain)
	}
