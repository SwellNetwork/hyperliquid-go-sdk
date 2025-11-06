package hyperliquid

import (
	"context"
	"fmt"
	"time"

	"resty.dev/v3"
)

type HTTPClientConfig struct {
	BaseURL string
	Timeout time.Duration
}
type HTTPClient struct {
	config HTTPClientConfig

	client *resty.Client
}

func NewHTTPClient(config HTTPClientConfig) *HTTPClient {
	return &HTTPClient{
		config: config,
		client: resty.New().SetBaseURL(config.BaseURL).SetTimeout(config.Timeout).SetHeader("Content-Type", "application/json"),
	}
}

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

func (c *HTTPClient) Meta(ctx context.Context, dex string) (*MetaResult, error) {
	body := map[string]any{
		"type": InfoTypeMeta,
	}

	if len(dex) > 0 {
		body["dex"] = dex
	}

	var result MetaResult

	resp, err := c.client.R().SetContext(ctx).SetBody(body).SetResult(&result).Post(PathInfo)
	if err != nil {
		return nil, fmt.Errorf("meta request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("meta request failed: status %d: %s", resp.StatusCode(), resp.String())
	}

	return &result, nil
}
