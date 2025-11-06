package hyperliquid

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
	FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error)
	Meta(ctx context.Context, dex string) (*MetaResult, error)
}

type PredictedFunding struct {
	Coin                 string
	Exchange             string
	FundingRates         float64
	NextFundingTime      int64
	FundingIntervalHours int64
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

type MetaResult struct {
	Universe []CoinMeta `json:"universe"`
}

type CoinMeta struct {
	SzDecimals    int64  `json:"szDecimals"`
	Name          string `json:"name"`
	MaxLeverage   int64  `json:"maxLeverage"`
	MarginTableId int64  `json:"marginTableId"`
}
