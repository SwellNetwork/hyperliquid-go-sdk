package hyperliquid

import (
	"fmt"

	"github.com/goccy/go-json"
)

type dispatcher func(find sharedSubscriptionFinder, msg wsMessage) error

func newChannelDispatcher[T subscriptable](channel wsChannel) dispatcher {
	return func(find sharedSubscriptionFinder, msg wsMessage) error {
		if msg.Channel != channel {
			return nil
		}

		var payload T
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal message: %v", err)
		}

		key := payload.Key()
		if key == "" {
			return nil
		}

		s, ok := find(key)
		if !ok || s == nil {
			return nil
		}

		s.dispatch(payload)
		return nil
	}
}
