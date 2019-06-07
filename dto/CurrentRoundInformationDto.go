package dto

import "time"

type RoundDto struct {
	RoundNumber                       int64
	TermNumber                        int64
	RoundId                           int64
	RealTimeMinerInformation          map[string]MinerInRoundDto
	ExtraBlockProducerOfPreviousRound string
}

type MinerInRoundDto struct {
	Order              int32
	ProducedTinyBlocks int32
	ExpectedMiningTime time.Time
	ActualMiningTimes  []time.Time
	InValue            string
	PreviousInValue    string
	OutValue           string
	ProducedBlocks     int64
	MissedBlocks       int64
}
