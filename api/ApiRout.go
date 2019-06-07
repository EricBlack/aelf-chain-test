package api

const (
	//chain route
	GetChainInformation          = "/api/blockChain/chainStatus"
	GetChainStatus               = "/api/blockChain/chainStatus"
	GetBlockHeight               = "/api/blockChain/blockHeight"
	CreateRawTransaction         = "/api/blockChain/rawTransaction"
	GetTransactionPoolStatus     = "/api/blockChain/transactionPoolStatus"
	GetBlockByHeight             = "/api/blockChain/blockByHeight?blockHeight=0}&includeTransactions=1}"
	GetBlockByHash               = "/api/blockChain/block?blockHash=0}&includeTransactions=1}"
	DeploySmartContract          = "/api/blockChain/broadcastTransaction"
	BroadcastTransaction         = "/api/blockChain/broadcastTransaction"
	BroadcastTransactions        = "/api/blockChain/broadcastTransactions"
	SendRawTransaction           = "/api/blockChain/sendRawTransaction"
	GetBlockState                = "/api/blockChain/blockState?blockHash=0}"
	Call                         = "/api/blockChain/call"
	GetContractFileDescriptorSet = "/api/blockChain/contractFileDescriptorSet?address=0}"
	GetTransactionResult         = "/api/blockChain/transactionResult?transactionId=0}"
	GetTransactionResults        = "/api/blockChain/transactionResults?blockHash=0}&offset=1}&limit=2}"

	//net route
	GetPeers   = "api/net/peers"
	AddPeer    = "/api/net/peer"
	RemovePeer = "/api/net/peer?address=0}"
)
