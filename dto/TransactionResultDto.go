package dto

type TransactionResultDto struct {
	TransactionId string
	Status string
	Logs []LogEventDto
	Bloom string
	BlockNumber int64
	BlockHash string
	Transaction TransactionDto
	ReadableReturnValue string
	Error string
}
