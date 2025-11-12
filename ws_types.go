package hyperliquid

import (
	"strings"

	"github.com/goccy/go-json"
)

type subscriptable interface {
	Key() string
}

type wsCommand struct {
	Method       string `json:"method"`
	Subscription any    `json:"subscription,omitempty"`
}

var (
	wsCommandPing        = wsCommand{Method: "ping"}
	wsCommandSubscribe   = func(subscription any) wsCommand { return wsCommand{Method: "subscribe", Subscription: subscription} }
	wsCommandUnsubscribe = func(subscription any) wsCommand { return wsCommand{Method: "unsubscribe", Subscription: subscription} }
)

type wsMessage struct {
	Channel wsChannel       `json:"channel"`
	Data    json.RawMessage `json:"data"`
}

type wsChannel string

const (
	ChannelTrades         wsChannel = "trades"
	ChannelActiveAssetCtx wsChannel = "activeAssetCtx"
	//ChannelL2Book         wsChannel = "l2Book"
	//ChannelCandle         wsChannel = "candle"
	//ChannelAllMids        wsChannel = "allMids"
	//ChannelNotification   wsChannel = "notification"
	//ChannelOrderUpdates   wsChannel = "orderUpdates"
	//ChannelUserFills      wsChannel = "userFills"
	//ChannelWebData2       wsChannel = "webData2"
	//ChannelBbo            wsChannel = "bbo"
	//ChannelSubResponse    wsChannel = "subscriptionResponse"
)

type Subscription struct {
	ID      string
	Payload any
	Close   func()
}

func key(args ...string) string {
	return strings.Join(args, ":")
}
