package hyperliquid_go_sdk

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
		client: resty.New().SetBaseURL(config.BaseURL).SetTimeout(config.Timeout),
	}
}

func (c *HTTPClient) PredictedFundings(ctx context.Context) ([]PredictedFunding, error) {
	var result predictedFundingsResult

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(map[string]any{"type": InfoTypePredictedFundings}).
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
