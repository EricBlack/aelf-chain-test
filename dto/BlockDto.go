package dto

type BlockDto struct {
	BlockHash string
	Header    BlockHeaderDto
	Body      BlockBodyDto
}
