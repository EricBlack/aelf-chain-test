package contracts

import (
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/api"
)

const(
	ElectionName = "AElf.ContractNames.Election"
	ProfitName = "AElf.ContractNames.Profit"
	VoteSystemName = "AElf.ContractNames.Vote"
	TokenName = "AElf.ContractNames.Token"
	TokenConverterName = "AElf.ContractNames.TokenConverter"
	FeeReceiverName = "AElf.ContractNames.FeeReceiver"
	ConsensusName = "AElf.ContractNames.Consensus"
	ParliamentName = "AElf.ContractsName.Parliament"
	CrossChainName = "AElf.ContractNames.CrossChain"
)

type BasicZero struct {
	ContractAddress string
}

func GetBasicZeroContract(chain api.Chain) BasicZero {
	status :=chain.GetChainStatus()
	contractAddress :=status.GenesisContractAddress

	return BasicZero{
		ContractAddress: contractAddress,
	}
}

func GetContractAddressByHashName(hashName string) string {
	return ""
}

func GetContractHashName(name string) (hash aelf.Hash) {
	hash = aelf.Hash{
		Value: []byte(name),
	}
	return
}
