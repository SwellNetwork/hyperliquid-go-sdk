package hyperliquid

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
	FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error)
	MetaAndAssetCtxs(ctx context.Context) (*MetaAndAssetCtx, error)
	SpotMetaAndAssetCtxs(ctx context.Context) (*SpotMetaAndAssetCtx, error)
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
