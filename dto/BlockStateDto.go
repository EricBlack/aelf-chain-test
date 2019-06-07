package dto

type BlockStateDto struct {
	BlockHash string
	PreviousHash string
	BlockHeight int64
	Changes map[string]string
}
