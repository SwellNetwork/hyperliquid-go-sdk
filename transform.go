package hyperliquid

import (
	"fmt"
	"strconv"
)

func transformPredictedFundingsResult(result predictedFundingsResult) ([]PredictedFunding, error) {
	var items []PredictedFunding
	for _, coin := range result {
		for _, exchange := range coin.Exchanges {
			rate := 0.0
			if exchange.Details.FundingRate != "" {
				parsed, err := strconv.ParseFloat(exchange.Details.FundingRate, 64)
				if err != nil {
					return nil, fmt.Errorf("parse funding rate %q for %s/%s: %w", exchange.Details.FundingRate, coin.Coin, exchange.Exchange, err)
				}
				rate = parsed
			}

			items = append(items, PredictedFunding{
				Coin:                 coin.Coin,
				Exchange:             exchange.Exchange,
				FundingRates:         rate,
				NextFundingTime:      exchange.Details.NextFundingTime,
				FundingIntervalHours: exchange.Details.FundingIntervalHours,
			})
		}
	}

	return items, nil
}

func transformMetaAndAssetCtxsResult(result metaAndAssetCtxsResult) ([]MetaAndAssetCtx, error) {
	if len(result.Meta.Universe) != len(result.Assets) {
		return nil, fmt.Errorf("meta and asset ctxs response mismatch: universe %d assetCtxs %d", len(result.Meta.Universe), len(result.Assets))
	}

	items := make([]MetaAndAssetCtx, len(result.Meta.Universe))
	for i := range result.Meta.Universe {
		uni := result.Meta.Universe[i]
		asset := result.Assets[i]
		item := &items[i]

		item.Universe.SzDecimals = uni.SzDecimals
		item.Universe.Name = uni.Name
		item.Universe.MaxLeverage = uni.MaxLeverage

		marginTable, ok := result.Meta.MarginTables[uni.MarginTableID]
		if ok {
			item.MarginTable.Description = marginTable.Description
			if len(marginTable.MarginTiers) > 0 {
				item.MarginTable.MarginTiers = make([]struct {
					LowerBound  string `json:"lowerBound"`
					MaxLeverage int    `json:"maxLeverage"`
				}, len(marginTable.MarginTiers))
				for j, tier := range marginTable.MarginTiers {
					item.MarginTable.MarginTiers[j].LowerBound = tier.LowerBound
					item.MarginTable.MarginTiers[j].MaxLeverage = tier.MaxLeverage
				}
			} else {
				item.MarginTable.MarginTiers = nil
			}
		} else {
			item.MarginTable.Description = ""
			item.MarginTable.MarginTiers = nil
		}

		item.AssetCtx.Funding = asset.Funding
		item.AssetCtx.OpenInterest = asset.OpenInterest
		item.AssetCtx.PrevDayPx = asset.PrevDayPx
		item.AssetCtx.DayNtlVlm = asset.DayNtlVlm
		item.AssetCtx.Premium = asset.Premium
		item.AssetCtx.OraclePx = asset.OraclePx
		item.AssetCtx.MarkPx = asset.MarkPx
		item.AssetCtx.MidPx = asset.MidPx
		item.AssetCtx.ImpactPxs = asset.ImpactPxs
		item.AssetCtx.DayBaseVlm = asset.DayBaseVlm
	}

	return items, nil
}

func transformSpotMetaAndAssetCtxsResult(result spotMetaAndAssetCtxsResult) ([]SpotMetaAndAssetCtx, error) {
	items := make([]SpotMetaAndAssetCtx, len(result.Meta.Universe))
	for i := range result.Meta.Universe {
		uni := result.Meta.Universe[i]
		asset := result.Assets[i]
		item := &items[i]

		item.Universe.Name = uni.Name
		item.Universe.Index = uni.Index
		item.Universe.IsCanonical = uni.IsCanonical

		item.AssetCtx.PrevDayPx = asset.PrevDayPx
		item.AssetCtx.DayNtlVlm = asset.DayNtlVlm
		item.AssetCtx.MarkPx = asset.MarkPx
		item.AssetCtx.MidPx = asset.MidPx
		item.AssetCtx.CirculatingSupply = asset.CirculatingSupply
		item.AssetCtx.Coin = asset.Coin
		item.AssetCtx.TotalSupply = asset.TotalSupply
		item.AssetCtx.DayBaseVlm = asset.DayBaseVlm
	}

	return items, nil
}
