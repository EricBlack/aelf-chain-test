package api

import (
	"ericshu.io/aelf-chain-test/dto"
	"ericshu.io/aelf-chain-test/utils"
	"encoding/json"
)

type Chain struct{
	HostUrl string
	Client utils.HttpClient
}

func GetChain(hostUrl string) (chain Chain) {
	return Chain {
		HostUrl: hostUrl,
		Client: utils.NewClient(hostUrl),
	}
}

func (chain Chain) GetBlockHeight() (result int64) {
	response :=chain.Client.Get(GetBlockHeight, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = -1
	}

	return
}

func (chain Chain) GetChainInformation() (result dto.GetChainInformationOutput){
	response :=chain.Client.Get(GetChainInformation, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.GetChainInformationOutput{}
	}

	return
}