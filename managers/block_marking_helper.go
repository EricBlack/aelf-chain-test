package managers

import (
	"encoding/hex"
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/api"
	"ericshu.io/aelf-chain-test/utils"
	"github.com/golang/protobuf/proto"

	"time"
)

var	Chain api.Chain
var	refBlockTime time.Time
var	cacheHeight int64
var	cacheHash string
var	chainId string

func InitBlockMarking(chain api.Chain) {
	Chain = chain
	refBlockTime = time.Now()
	cacheHeight = Chain.GetBlockHeight()
	cacheHash = Chain.GetBlockByHeight(cacheHeight, false).BlockHash
	chainId = Chain.GetChainStatus().ChainId
}

func ConvertTransactionRawTxString(transaction aelf.Transaction) (rawString string) {
	bytes, _ := proto.Marshal(&transaction)
	rawString = hex.EncodeToString(bytes)

	return
}

func AddBlockReference(transaction aelf.Transaction) aelf.Transaction {
	height := cacheHeight
	if height ==0 || time.Now().Sub(refBlockTime).Seconds() > 30 {
		cacheHeight = Chain.GetBlockHeight()
		cacheHash =  Chain.GetBlockByHeight(cacheHeight, false).BlockHash
		refBlockTime = time.Now()
	}
	prefix, err := utils.FromHexString(cacheHash)
	if err != nil {
		return transaction
	}

	transaction.RefBlockNumber = cacheHeight
	transaction.RefBlockPrefix = prefix[0:3]

	return transaction
}
