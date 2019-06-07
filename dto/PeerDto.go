package dto

type PeerDto struct {
	IpAddress       string
	ProtocolVersion int
	ConnectionTime  int64
	Inbound         bool
	StartHeight     int64
	RequestMetrics  []RequestMetric
}
