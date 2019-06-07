package dto

type CreateRawTransactionInput struct {
	/// <summary>
	/// from address
	/// </summary>
	From string

	/// <summary>
	/// to address
	/// </summary>
	To string

	/// <summary>
	/// refer block height
	/// </summary>
	RefBlockNumber int64

	/// <summary>
	/// refer block hash
	/// </summary>
	RefBlockHash string

	/// <summary>
	/// contract method name
	/// </summary>
	MethodName string

	/// <summary>
	/// contract method parameters
	/// </summary>
	Params string
}
