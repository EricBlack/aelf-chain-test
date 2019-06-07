package dto

type SendRawTransactionInput struct {
	/// <summary>
	/// raw transaction
	/// </summary>
	Transaction string

	/// <summary>
	/// signature
	/// </summary>
	Signature string

	/// <summary>
	/// return transaction detail or not
	/// </summary>
	ReturnTransaction bool
}
