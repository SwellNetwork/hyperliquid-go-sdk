package hyperliquid

type TradesSubscriptionParams struct {
	Coin string
}

type Trade struct {
	Coin  string   `json:"coin"`
	Side  string   `json:"side"`
	Px    string   `json:"px"`
	Sz    string   `json:"sz"`
	Time  int64    `json:"time"`
	Hash  string   `json:"hash"`
	Tid   int64    `json:"tid"`
	Users []string `json:"users"`
}

type Trades []Trade

func (t Trades) Key() string {
	if len(t) == 0 {
		return ""
	}
	return genKey(string(ChannelTrades), t[0].Coin)
}

type tradeSubscriptionPayload struct {
	Type wsChannel `json:"type"`
	Coin string    `json:"coin"`
}

func (p tradeSubscriptionPayload) Key() string {
	return genKey(string(ChannelTrades), p.Coin)
}

func (c *WSClient) Trades(
	params TradesSubscriptionParams,
	callback func([]Trade, error),
) (*Subscription, error) {
	remotePayload := tradeSubscriptionPayload{
		Type: ChannelTrades,
		Coin: params.Coin,
	}

	return subscribeTyped[Trades](c, remotePayload, func(trades Trades, err error) {
		if err != nil {
			callback(nil, err)
			return
		}

		callback([]Trade(trades), nil)
	})
}
