package hyperliquid

import "time"

type HTTPClientConfig struct {
	BaseURL string
	Timeout time.Duration
}

type WSClientConfig struct {
	BaseURL           string
	PingInterval      time.Duration
	ReconnectAttempts int
}

func DefaultMainnetHTTPClientConfig() HTTPClientConfig {
	return HTTPClientConfig{
		BaseURL: "https://api.hyperliquid.xyz",
		Timeout: 3 * time.Second,
	}
}

func DefaultTestnetHTTPClientConfig() HTTPClientConfig {
	return HTTPClientConfig{
		BaseURL: "https://api.hyperliquid-testnet.xyz",
		Timeout: 3 * time.Second,
	}
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
