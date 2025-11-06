package hyperliquid_go_sdk

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
	FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error)
}

type PredictedFunding struct {
	Coin                 string
	Exchange             string
	FundingRates         float64
	NextFundingTime      int64
	FundingIntervalHours int
}

type FundingHistoryParams struct {
	Coin      string
	StartTime int64
	EndTime   int64
}

type FundingHistory struct {
	Coin        string `json:"coin"`
	FundingRate string `json:"fundingRate"`
	Premium     string `json:"premium"`
	Time        int64  `json:"time"`
}
