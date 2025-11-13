package hyperliquid

import "time"

type HTTPClientConfig struct {
	BaseURL string
	Timeout time.Duration
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
