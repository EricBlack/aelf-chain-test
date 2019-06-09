package main

import (
	"ericshu.io/aelf-chain-test/api"
	"fmt"
	"log"
)

var chain api.Chain

func init() {
	chain = api.GetChain("http://127.0.0.1:8000")
}

func main(){
	chainInfo := chain.GetChainStatus()
	fmt.Println("ChainId: %s, Zero Contract: %s", chainInfo.ChainId, chainInfo.GenesisContractAddress)

	height := chain.GetBlockHeight()
	fmt.Println("Block height: ", height)

	var i int64
	for i=1; i<=height; i++ {
		block := chain.GetBlockByHeight(i, true)
		log.Printf(fmt.Sprintf("Height: %d, Block info: %s", i, block))
	}
}
