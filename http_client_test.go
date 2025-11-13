//go:build integration
// +build integration

package hyperliquid

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestHTTPClientIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPClientIntegrationTestSuite))
}

type HTTPClientIntegrationTestSuite struct {
	suite.Suite

	client *HTTPClient
}

func (ts *HTTPClientIntegrationTestSuite) SetupSuite() {
	config := HTTPClientConfig{
		BaseURL: "https://api.hyperliquid.xyz",
		Timeout: 3 * time.Second,
	}

	ts.client = NewHTTPClient(config)
}

func (ts *HTTPClientIntegrationTestSuite) TestPredictedFundings() {
	ctx := context.Background()

	result, err := ts.client.PredictedFundings(ctx)

	ts.T().Log("predicted fundings", result)

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestFundingHistory() {
	ctx := context.Background()

	params := FundingHistoryParams{
		Coin:      "ETH",
		StartTime: time.Now().Add(-1 * time.Hour).UnixMilli(),
		EndTime:   time.Now().UnixMilli(),
	}

	result, err := ts.client.FundingHistory(ctx, params)
	ts.T().Log("funding history", result)

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestMetaAndAssetCtxs() {
	ctx := context.Background()

	result, err := ts.client.MetaAndAssetCtxs(ctx)
	ts.T().Log("meta and asset contexts", len(result.Meta.Universe))

	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *HTTPClientIntegrationTestSuite) TestSpotMetaAndAssetCtxsResult() {
	ctx := context.Background()
	result, err := ts.client.SpotMetaAndAssetCtxs(ctx)
	ts.T().Log("spot meta and asset contexts", len(result.Ctxs))

	ts.NoError(err)
	ts.NotNil(result)
}
