package hyperliquid_go_sdk

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
}

type PredictedFunding struct {
	Coin                 string
	Exchange             string
	FundingRates         float64
	NextFundingTime      int64
	FundingIntervalHours int
}
