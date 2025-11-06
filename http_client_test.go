package hyperliquid_go_sdk

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

	ts.T().Log("result", result)

	ts.NoError(err)
	ts.NotNil(result)
}
