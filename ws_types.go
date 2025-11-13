package hyperliquid

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
)

type sharedSubscriptionFinder func(string) (*sharedSubscription, bool)
type callback func(any)

type subscriptable interface {
	Key() string
}

func subscribeTyped[T any](
	c *WSClient,
	payload subscriptable,
	callback func(T, error),
) (*Subscription, error) {
	if callback == nil {
		return nil, fmt.Errorf("callback cannot be nil")
	}

	var zero T

	return c.subscribe(payload, func(msg any) {
		typed, ok := msg.(T)
		if !ok {
			callback(zero, fmt.Errorf("invalid message type: %T", msg))
			return
		}

		callback(typed, nil)
	})
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
)

type Subscription struct {
	ID      string
	Payload any
	Close   func()
}

func genKey(args ...string) string {
	return strings.Join(args, ":")
}
