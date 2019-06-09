package api

import (
	"encoding/json"
	"ericshu.io/aelf-chain-test/dto"
	"ericshu.io/aelf-chain-test/utils"
)

type Chain struct {
	HostUrl string
	Client  utils.HttpClient
}

func GetChain(hostUrl string) (chain Chain) {
	chain = Chain{
		HostUrl: hostUrl,
		Client:  utils.NewClient(hostUrl),
	}
	chain.Client.SetDefaultHeaders()

	return
}

func (chain Chain) GetBlockHeight() (result int64) {
	response := chain.Client.Get(GetBlockHeight, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = -1
	}

	return
}

func (chain Chain) GetChainStatus() (result dto.GetChainInformationOutput) {
	response := chain.Client.Get(GetChainStatus, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.GetChainInformationOutput{}
	}

	return
}

func (chain Chain) CreateRawTransaction(inputMessage dto.CreateRawTransactionInput) (result dto.CreateRawTransactionOutput) {
	response := chain.Client.Post(CreateRawTransaction, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.CreateRawTransactionOutput{}
	}

	return
}

func (chain Chain) BroadcastTransaction(inputMessage dto.BroadcastTransactionInput) (result dto.BroadcastTransactionOutput) {
	response := chain.Client.Post(BroadcastTransaction, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.BroadcastTransactionOutput{}
	}

	return
}

func (chain Chain) BroadcastTransactions(inputMessage dto.BroadcastTransactionsInput) (result []string) {
	response := chain.Client.Post(BroadcastTransactions, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = []string{}
	}

	return
}

func (chain Chain) GetTransactionPoolStatus() (result dto.GetTransactionPoolStatusOutput) {
	response := chain.Client.Get(GetTransactionPoolStatus, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.GetTransactionPoolStatusOutput{}
	}

	return
}

func (chain Chain) GetBlockByHeight(height int64, includeTransactions bool) (result dto.BlockDto) {
	params := map[string]interface{}{
		"height":              height,
		"includeTransactions": includeTransactions,
	}

	response := chain.Client.Get(GetBlockByHeight, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.BlockDto{}
	}

	return
}

func (chain Chain) GetBlockByHash(hash string, includeTransactions bool) (result dto.BlockDto) {
	params := map[string]interface{}{
		"blockHash":           hash,
		"includeTransactions": includeTransactions,
	}

	response := chain.Client.Get(GetBlockByHash, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.BlockDto{}
	}

	return
}

func (chain Chain) GetBlockState(blockHash string) (result dto.BlockStateDto) {
	params := map[string]interface{}{
		"blockHash": blockHash,
	}

	response := chain.Client.Get(GetBlockState, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.BlockStateDto{}
	}

	return
}

func (chain Chain) Call(inputMessage dto.CallInput) (result []byte) {
	result = []byte{}

	return
}

func (chain Chain) GetContractFileDescriptorSet(address string) (result []byte) {
	params := map[string]interface{}{
		"address": address,
	}

	response := chain.Client.Get(GetContractFileDescriptorSet, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = []byte{}
	}

	result = []byte(response)
	return
}

func (chain Chain) GetTransactionResult(transactionId string) (result dto.TransactionResultDto) {
	params := map[string]interface{}{
		"transactionId": transactionId,
	}

	response := chain.Client.Get(GetTransactionResult, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = dto.TransactionResultDto{}
	}

	return
}

func (chain Chain) GetTransactionResults(blockHash string, offset, limit int32) (result []dto.TransactionResultDto) {
	params := map[string]interface{}{
		"blockHash": blockHash,
		"offset": offset,
		"limit": limit,
	}

	response := chain.Client.Get(GetTransactionResults, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = []dto.TransactionResultDto{}
	}

	return
}