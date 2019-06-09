package managers

import (
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/api"
	"ericshu.io/aelf-chain-test/dto"
	"ericshu.io/aelf-chain-test/utils"
	"github.com/golang/protobuf/proto"
)

type ContractFactory struct {
	Chain api.Chain
	Owner AddressInfo
	ContractAddress string
}

func GetContractFactory(chain api.Chain, info AddressInfo, contract string) (cf ContractFactory) {
	cf = ContractFactory{
		Chain: chain,
		Owner: info,
		ContractAddress: contract,
	}

	InitBlockMarking(chain)

	return
}

func (cf ContractFactory) SendTransaction(method string, inputMessage proto.Message) (result dto.BroadcastTransactionOutput){
	transaction := cf.GenerateTransaction(method, inputMessage)
	data := GetTransactionBytes(transaction)
	rawString := utils.FromBytesToHex(data)

	result = cf.Chain.BroadcastTransaction(dto.BroadcastTransactionInput{
		RawTransaction: rawString,
	})

	return
}

func  (cf ContractFactory) CallTransaction(method string, inputMessage proto.Message) (result []byte){
	transaction := cf.GenerateTransaction(method, inputMessage)
	data := GetTransactionBytes(transaction)
	rawString := utils.FromBytesToHex(data)

	result = cf.Chain.Call(dto.CallInput{
		RawTransaction: rawString,
	})

	return
}

func  (cf ContractFactory) GenerateTransaction(method string, inputMessage proto.Message) (transaction aelf.Transaction){
	params , _ := proto.Marshal(inputMessage)
	toData, _ := utils.FromHexString(cf.ContractAddress)
	transaction = aelf.Transaction{
		From: &aelf.Address{
			Value: cf.Owner.PublicKey,
		},
		To: &aelf.Address{
			Value: toData,
		},
		MethodName: method,
		Params: params,
	}

	//add block reference
	transaction = AddBlockReference(transaction)

	//sign transaction
	transaction = cf.Owner.SignTransaction(transaction)

	return
}
