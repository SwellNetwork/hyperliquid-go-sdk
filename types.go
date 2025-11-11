package hyperliquid

type (
	AssetInfo struct {
		Name          string `json:"name"`
		SzDecimals    int    `json:"szDecimals"`
		MaxLeverage   int    `json:"maxLeverage"`
		MarginTableId int    `json:"marginTableId"`
		OnlyIsolated  bool   `json:"onlyIsolated,omitempty"`
		IsDelisted    bool   `json:"isDelisted,omitempty"`
	}
	AssetCtx struct {
		Funding      string   `json:"funding"`
		OpenInterest string   `json:"openInterest"`
		PrevDayPx    string   `json:"prevDayPx"`
		DayNtlVlm    string   `json:"dayNtlVlm"`
		Premium      string   `json:"premium"`
		OraclePx     string   `json:"oraclePx"`
		MarkPx       string   `json:"markPx"`
		MidPx        string   `json:"midPx,omitempty"`
		ImpactPxs    []string `json:"impactPxs"`
		DayBaseVlm   string   `json:"dayBaseVlm,omitempty"`
	}

	MarginTier struct {
		LowerBound  string `json:"lowerBound"`
		MaxLeverage int    `json:"maxLeverage"`
	}

	MarginTable struct {
		ID          int
		Description string       `json:"description"`
		MarginTiers []MarginTier `json:"marginTiers"`
	}

	Meta struct {
		Universe     []AssetInfo   `json:"universe"`
		MarginTables []MarginTable `json:"marginTables"`
	}

	MetaAndAssetCtx struct {
		Meta
		Ctxs []AssetCtx
	}
)

type (
	SpotAssetInfo struct {
		Name        string `json:"name"`
		Tokens      []int  `json:"tokens"`
		Index       int    `json:"index"`
		IsCanonical bool   `json:"isCanonical"`
	}

	EvmContract struct {
		Address             string `json:"address"`
		EvmExtraWeiDecimals int    `json:"evm_extra_wei_decimals"`
	}

	SpotTokenInfo struct {
		Name        string       `json:"name"`
		SzDecimals  int          `json:"szDecimals"`
		WeiDecimals int          `json:"weiDecimals"`
		Index       int          `json:"index"`
		TokenID     string       `json:"tokenId"`
		IsCanonical bool         `json:"isCanonical"`
		EvmContract *EvmContract `json:"evmContract"`
		FullName    *string      `json:"fullName"`
	}
	SpotMeta struct {
		Universe []SpotAssetInfo `json:"universe"`
		Tokens   []SpotTokenInfo `json:"tokens"`
	}

	SpotAssetCtx struct {
		DayNtlVlm         string `json:"dayNtlVlm"`
		MarkPx            string `json:"markPx"`
		MidPx             string `json:"midPx"`
		PrevDayPx         string `json:"prevDayPx"`
		CirculatingSupply string `json:"circulatingSupply"`
		Coin              string `json:"coin"`
		TotalSupply       string `json:"totalSupply"`
		DayBaseVlm        string `json:"dayBaseVlm"`
	}

	SpotMetaAndAssetCtx struct {
		Meta SpotMeta
		Ctxs []SpotAssetCtx
	}
)

type (
	FundingHistory struct {
		Coin        string `json:"coin"`
		FundingRate string `json:"fundingRate"`
		Premium     string `json:"premium"`
		Time        int64  `json:"time"`
	}
)
