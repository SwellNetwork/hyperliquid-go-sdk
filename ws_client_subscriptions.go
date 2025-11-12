package hyperliquid

import (
	"fmt"

	"github.com/google/uuid"
)

func (c *WSClient) subscribe(
	payload subscriptable,
	callback func(any),
) (*Subscription, error) {
	if callback == nil {
		return nil, fmt.Errorf("callback cannot be nil")
	}

	c.mu.Lock()

	pkey := payload.Key()
	s, exists := c.sharedSubscriptionByKey[pkey]
	if !exists {
		s = newSubscriber(
			pkey,
			payload,
			func(p subscriptable) {
				if err := c.writeJSON(wsCommandSubscribe(payload)); err != nil {
					c.logger.Errorf("failed to subscribe: %v", err)
				}
			},
			func(p subscriptable) {
				c.mu.Lock()
				defer c.mu.Unlock()
				delete(c.sharedSubscriptionByKey, pkey)

				if err := c.writeJSON(wsCommandUnsubscribe(payload)); err != nil {
					c.logger.Errorf("failed to unsubscribe: %v", err)
				}
			},
		)

		c.sharedSubscriptionByKey[pkey] = s
	}

	c.mu.Unlock()

	subscriberID := uuid.New().String()
	s.addSubscriber(subscriberID, callback)

	return &Subscription{
		ID: subscriberID,
		Close: func() {
			s.removeSubscriber(subscriberID)
		},
	}, nil
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

func (c *WSClient) resubscribeAll() error {
	c.mu.RLock()
	payloads := make([]subscriptable, 0, len(c.sharedSubscriptionByKey))
	for _, s := range c.sharedSubscriptionByKey {
		if payload, ok := s.activePayload(); ok {
			payloads = append(payloads, payload)
		}
	}
	c.mu.RUnlock()

	for _, payload := range payloads {
		if err := c.writeJSON(wsCommandSubscribe(payload)); err != nil {
			return err
		}
	}

	return nil
}
