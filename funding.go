package hyperliquid

import (
	"context"
	"fmt"
)

func (c *HTTPClient) PredictedFundings(ctx context.Context) ([]PredictedFunding, error) {
	body := map[string]any{"type": InfoTypePredictedFundings}

	var result predictedFundingsResult

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&result).
		Post(PathInfo)
	if err != nil {
		return nil, fmt.Errorf("predicted fundings request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("predicted fundings request failed: status %d: %s", resp.StatusCode(), resp.String())
	}

	return transformPredictedFundingsResult(result)
}

func (c *HTTPClient) FundingHistory(ctx context.Context, params FundingHistoryParams) ([]FundingHistory, error) {
	body := map[string]any{
		"type":      InfoTypeFundingHistory,
		"coin":      params.Coin,
		"startTime": params.StartTime,
	}
	if params.EndTime > 0 {
		body["endTime"] = params.EndTime
	}

	var result []FundingHistory

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&result).
		Post(PathInfo)
	if err != nil {
		return nil, fmt.Errorf("funding history request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("funding history request failed: status %d: %s", resp.StatusCode(), resp.String())
	}

	return result, nil
}
