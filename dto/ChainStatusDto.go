package dto

type ChainStatusDto struct {
	ChainId                     string
	Branches                    map[string]int64
	NotLinkedBlocks             []NotLinkedBlockDto
	LongestChainHeight          int64
	LongestChainHash            string
	GenesisBlockHash            string
	GenesisContractAddress      string
	LastIrreversibleBlockHash   string
	LastIrreversibleBlockHeight int64
	BestChainHash               string
	BestChainHeight             int64
}
