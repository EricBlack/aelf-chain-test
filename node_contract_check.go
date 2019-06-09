package main

import (
	"encoding/json"
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/api"
	"ericshu.io/aelf-chain-test/managers"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func init() {
	chain = api.GetChain("http://127.0.0.1:8000")
}

func main(){
	contract := chain.GetChainStatus().GenesisContractAddress
	info := managers.GenerateNewAddress()
	contractFactory := managers.GetContractFactory(chain, info, contract)

	result :=contractFactory.CallTransaction("GetContractInfo", &empty.Empty{})
	address := aelf.Address{}
	json.Unmarshal(result, &address)
	log.Println("Address info: ",address.Value)
}
