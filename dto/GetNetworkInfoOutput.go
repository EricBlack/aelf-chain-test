package dto

type GetNetworkInfoOutput struct {
	/// <summary>
	/// node version
	/// </summary>
	Version string

	/// <summary>
	/// network protocol version
	/// </summary>
	ProtocolVersion int

	/// <summary>
	/// total number of open connections between this node and other nodes
	/// </summary>
	Connections int
}
