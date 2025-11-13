package hyperliquid

type ActiveAssetCtxSubscriptionParams struct {
	Coin string
}

type (
	ActiveAssetCtx struct {
		Coin string         `json:"coin"`
		Ctx  SharedAssetCtx `json:"ctx"`
	}

	SharedAssetCtx struct {
		DayNtlVlm float64 `json:"dayNtlVlm,string"`
		PrevDayPx float64 `json:"prevDayPx,string"`
		MarkPx    float64 `json:"markPx,string"`
		MidPx     float64 `json:"midPx,string"`

		Funding      float64 `json:"funding,string,omitempty"`
		OpenInterest float64 `json:"openInterest,string,omitempty"`
		OraclePx     float64 `json:"oraclePx,string,omitempty"`

		CirculatingSupply float64 `json:"circulatingSupply,string,omitempty"`
	}
)

func (a ActiveAssetCtx) Key() string {
	return genKey(string(ChannelActiveAssetCtx), a.Coin)
}

type activeAssetCtxSubscriptionPayload struct {
	Type wsChannel `json:"type"`
	Coin string    `json:"coin"`
}

func (p activeAssetCtxSubscriptionPayload) Key() string {
	return genKey(string(ChannelActiveAssetCtx), p.Coin)
}

func (c *WSClient) ActiveAssetCtx(
	params ActiveAssetCtxSubscriptionParams,
	callback func(ActiveAssetCtx, error),
) (*Subscription, error) {
	remotePayload := activeAssetCtxSubscriptionPayload{
		Type: ChannelActiveAssetCtx,
		Coin: params.Coin,
	}

	return subscribeTyped[ActiveAssetCtx](c, remotePayload, callback)
}
