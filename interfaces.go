package hyperliquid

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
	FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error)
	MetaAndAssetCtxs(ctx context.Context) (*MetaAndAssetCtx, error)
	
	SpotMetaAndAssetCtxs(ctx context.Context) (*SpotMetaAndAssetCtx, error)
}
