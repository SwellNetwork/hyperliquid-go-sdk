//go:build integration
// +build integration

package hyperliquid

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestWSClientIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(WSClientIntegrationTestSuite))
}

type WSClientIntegrationTestSuite struct {
	suite.Suite

	client *WSClient
}

func (ts *WSClientIntegrationTestSuite) SetupTest() {
	ts.client = NewWSClient(
		DefaultMainnetWSClientConfig(),
		WithWSClientDebug(true),
	)
}

func (ts *WSClientIntegrationTestSuite) TestTrade() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ts.T().Log("Connecting to websocket")
	if err := ts.client.Connect(ctx); err != nil {
		ts.T().Fatalf("Failed to connect: %v", err)
	}
	defer ts.client.Close()

	sub, err := ts.client.Trades(
		TradesSubscriptionParams{Coin: "ETH"},
		func(trades []Trade, err error) {
			if err != nil {
				ts.T().Fatalf("Failed to receive trades: %v", err)
			}
			ts.T().Logf("Received a set of trades: %+v", len(trades))
		},
	)

	if err != nil {
		ts.T().Fatalf("Failed to subscribe to trades: %v", err)
	}

	ts.T().Log("Subscribed to trades")

	defer sub.Close()

	<-time.After(time.Second * 5)
	ts.T().Log("Unsubscribing from trades")
	sub.Close()
}

func (ts *WSClientIntegrationTestSuite) TestActiveAssetCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ts.T().Log("Connecting to websocket")
	if err := ts.client.Connect(ctx); err != nil {
		ts.T().Fatalf("Failed to connect: %v", err)
	}
	defer ts.client.Close()

	sub, err := ts.client.ActiveAssetCtx(
		ActiveAssetCtxSubscriptionParams{Coin: "ETH"},
		func(act ActiveAssetCtx, err error) {
			if err != nil {
				ts.T().Fatalf("Failed to receive active asset context: %v", err)
			}

			ts.T().Logf("Received a set of active asset context: %+v", act)
		},
	)
	if err != nil {
		ts.T().Fatalf("Failed to subscribe to active asset context: %v", err)
	}
	ts.T().Log("Subscribed to active asset context")
	defer sub.Close()

	<-time.After(time.Second * 5)
	ts.T().Log("Unsubscribing from active asset context")
}
