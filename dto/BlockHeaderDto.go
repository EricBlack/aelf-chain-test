package dto

import "time"

type BlockHeaderDto struct {
	PreviousBlockHash            string
	MerkleTreeRootOfTransactions string
	MerkleTreeRootOfWorldState   string
	Extra                        string
	Height                       int64
	Time                         time.Time
	ChainId                      string
	Bloom                        string
	SignerPubkey                 string
}
