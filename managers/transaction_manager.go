package managers

import (
	"ericshu.io/aelf-chain-test/api"
	"time"
)

type TransactionManager struct {
	Chain api.Chain
	refBlockTime time.Time
	cacheHeight int64
	cacheHash string
}
