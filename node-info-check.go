package main

import (
	"ericshu.io/aelf-chain-test/api"
	"fmt"
)

var chain api.Chain

func init() {
	chain = api.GetChain("http://127.0.0.1:8000")
}

func main(){
	height := chain.GetBlockHeight()
	fmt.Println("Block height: ", height)

	chainInfo := chain.GetChainStatus()
	fmt.Println("ChainId:", chainInfo.ChainId)
}
