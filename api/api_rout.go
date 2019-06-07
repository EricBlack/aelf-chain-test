package api

const (
	//chain route
	GetChainStatus               = "/api/blockChain/chainStatus"
	GetBlockHeight               = "/api/blockChain/blockHeight"
	CreateRawTransaction         = "/api/blockChain/rawTransaction"
	GetTransactionPoolStatus     = "/api/blockChain/transactionPoolStatus"
	GetBlockByHeight             = "/api/blockChain/blockByHeight"
	GetBlockByHash               = "/api/blockChain/block"
	BroadcastTransaction         = "/api/blockChain/broadcastTransaction"
	BroadcastTransactions        = "/api/blockChain/broadcastTransactions"
	SendRawTransaction           = "/api/blockChain/sendRawTransaction"
	GetBlockState                = "/api/blockChain/blockState"
	Call                         = "/api/blockChain/call"
	GetContractFileDescriptorSet = "/api/blockChain/contractFileDescriptorSet"
	GetTransactionResult         = "/api/blockChain/transactionResult"
	GetTransactionResults        = "/api/blockChain/transactionResults"

	//net route
	GetPeers   = "api/net/peers"
	AddPeer    = "/api/net/peer"
	RemovePeer = "/api/net/peer"
)
