package managers

import (
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/cryptography"
	"github.com/golang/protobuf/proto"
)

type TransactionManager struct {
}

func GetHash(signature []byte) (hash aelf.Hash) {
	hash = aelf.Hash{
		Value: cryptography.CalculateHash(signature),
	}
	return
}

func GetHashBytes(signature []byte) []byte {
	return cryptography.CalculateHash(signature)
}

func GetSignatureData(transaction aelf.Transaction) []byte {
	data , _ := proto.Marshal(&transaction)

	return data
}


