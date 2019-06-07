package dto

import "github.com/golang/protobuf/ptypes/timestamp"

type RequestMetric struct {
	RoundTripTime int64
	MethodName    string
	Info          string
	RequestTime   timestamp.Timestamp
}
