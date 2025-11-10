package hyperliquid

import "context"

type InfoClient interface {
	PredictedFundings(ctx context.Context) ([]PredictedFunding, error)
	FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error)
	MetaAndAssetCtxs(ctx context.Context) ([]MetaAndAssetCtx, error)
	SpotMetaAndAssetCtxs(ctx context.Context) ([]SpotMetaAndAssetCtx, error)
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

type MetaAndAssetCtx struct {
	Universe struct {
		SzDecimals  int    `json:"szDecimals"`
		Name        string `json:"name"`
		MaxLeverage int    `json:"maxLeverage"`
	} `json:"universe"`
	MarginTable struct {
		Description string `json:"description"`
		MarginTiers []struct {
			LowerBound  string `json:"lowerBound"`
			MaxLeverage int    `json:"maxLeverage"`
		} `json:"marginTiers"`
	} `json:"marginTable,omitempty"`
	AssetCtx struct {
		Funding      string   `json:"funding"`
		OpenInterest string   `json:"openInterest"`
		PrevDayPx    string   `json:"prevDayPx"`
		DayNtlVlm    string   `json:"dayNtlVlm"`
		Premium      string   `json:"premium"`
		OraclePx     string   `json:"oraclePx"`
		MarkPx       string   `json:"markPx"`
		MidPx        string   `json:"midPx"`
		ImpactPxs    []string `json:"impactPxs"`
		DayBaseVlm   string   `json:"dayBaseVlm"`
	} `json:"assetCtx"`
}

type SpotMetaAndAssetCtx struct {
	Universe struct {
		Name        string `json:"name"`
		Index       int    `json:"index"`
		IsCanonical bool   `json:"isCanonical"`
	}
	AssetCtx struct {
		PrevDayPx         string `json:"prevDayPx"`
		DayNtlVlm         string `json:"dayNtlVlm"`
		MarkPx            string `json:"markPx"`
		MidPx             string `json:"midPx"`
		CirculatingSupply string `json:"circulatingSupply"`
		Coin              string `json:"coin"`
		TotalSupply       string `json:"totalSupply"`
		DayBaseVlm        string `json:"dayBaseVlm"`
	} `json:"assetCtx"`
}
