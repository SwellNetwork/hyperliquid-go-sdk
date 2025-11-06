package hyperliquid_go_sdk

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
