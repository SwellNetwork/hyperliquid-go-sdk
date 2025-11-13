package hyperliquid

import "time"

type WSClientConfig struct {
	BaseURL           string
	PingInterval      time.Duration
	ReconnectAttempts int
}

func DefaultMainnetWSClientConfig() WSClientConfig {
	return WSClientConfig{
		BaseURL:           "wss://api.hyperliquid.xyz/ws",
		PingInterval:      30 * time.Second,
		ReconnectAttempts: 5,
	}
}

func DefaultTestnetWSClientConfig() WSClientConfig {
	return WSClientConfig{
		BaseURL:           "wss://api.hyperliquid-testnet.xyz/ws",
		PingInterval:      30 * time.Second,
		ReconnectAttempts: 5,
	}
}
