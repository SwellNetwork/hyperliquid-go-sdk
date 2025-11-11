package hyperliquid

import (
	"fmt"
	"sort"
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

func transformMetaAndAssetCtxsResult(result metaAndAssetCtxsResult) (*MetaAndAssetCtx, error) {
	if len(result.Meta.Universe) != len(result.Assets) {
		return nil, fmt.Errorf("meta and asset ctxs response mismatch: universe %d assetCtxs %d", len(result.Meta.Universe), len(result.Assets))
	}

	meta := Meta{
		Universe: make([]AssetInfo, len(result.Meta.Universe)),
	}

	for i, uni := range result.Meta.Universe {
		meta.Universe[i] = AssetInfo{
			Name:          uni.Name,
			SzDecimals:    uni.SzDecimals,
			MaxLeverage:   uni.MaxLeverage,
			MarginTableId: uni.MarginTableID,
			OnlyIsolated:  uni.OnlyIsolated,
			IsDelisted:    uni.IsDelisted,
		}
	}

	if len(result.Meta.MarginTables) > 0 {
		meta.MarginTables = make([]MarginTable, 0, len(result.Meta.MarginTables))

		tableIDs := make([]int, 0, len(result.Meta.MarginTables))
		for id := range result.Meta.MarginTables {
			tableIDs = append(tableIDs, id)
		}
		if len(tableIDs) > 1 {
			sort.Ints(tableIDs)
		}

		for _, id := range tableIDs {
			marginTable := result.Meta.MarginTables[id]
			tierCopy := make([]MarginTier, len(marginTable.MarginTiers))
			for i, tier := range marginTable.MarginTiers {
				tierCopy[i] = MarginTier(tier)
			}

			meta.MarginTables = append(meta.MarginTables, MarginTable{
				ID:          id,
				Description: marginTable.Description,
				MarginTiers: tierCopy,
			})
		}
	}

	ctxs := make([]AssetCtx, len(result.Assets))
	for i, asset := range result.Assets {
		ctxs[i] = AssetCtx(asset)
	}

	return &MetaAndAssetCtx{
		Meta: meta,
		Ctxs: ctxs,
	}, nil
}

func transformSpotMetaAndAssetCtxsResult(result spotMetaAndAssetCtxsResult) (*SpotMetaAndAssetCtx, error) {
	meta := SpotMeta{
		Universe: make([]SpotAssetInfo, len(result.Meta.Universe)),
		Tokens:   result.Meta.Tokens,
	}

	for i, uni := range result.Meta.Universe {
		meta.Universe[i] = SpotAssetInfo{
			Tokens:      append([]int(nil), uni.Tokens...),
			Name:        uni.Name,
			Index:       uni.Index,
			IsCanonical: uni.IsCanonical,
		}
	}

	ctxs := make([]SpotAssetCtx, len(result.Assets))
	for i, asset := range result.Assets {
		ctxs[i] = SpotAssetCtx(asset)
	}

	return &SpotMetaAndAssetCtx{
		Meta: meta,
		Ctxs: ctxs,
	}, nil
}
