package hyperliquid

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sonirico/vago/lol"
)

type WSClient struct {
	config WSClientConfig

	conn *websocket.Conn

	sharedSubscriptionByKey map[string]*sharedSubscription
	dispatcherByChannel     map[wsChannel]dispatcher

	done        chan struct{}
	reconnectCh chan struct{}
	mu          sync.RWMutex
	writeMu     sync.Mutex
	closeOnce   sync.Once
	debug       bool

	logger lol.Logger
}

type WSClientOption func(*WSClient)

func WithWSClientLogger(logger lol.Logger) WSClientOption {
	return func(c *WSClient) {
		l := logger
		if l == nil {
			l = lol.NewZerolog()
		}
		c.logger = l
	}
}

func WithWSClientDebug(debug bool) WSClientOption {
	return func(c *WSClient) {
		c.debug = debug
	}
}

func NewWSClient(config WSClientConfig, opts ...WSClientOption) *WSClient {
	client := &WSClient{
		config: config,

		sharedSubscriptionByKey: make(map[string]*sharedSubscription),
		dispatcherByChannel: map[wsChannel]dispatcher{
			ChannelTrades:         newChannelDispatcher[Trades](ChannelTrades),
			ChannelActiveAssetCtx: newChannelDispatcher[ActiveAssetCtx](ChannelActiveAssetCtx),
		},

		done:        make(chan struct{}),
		reconnectCh: make(chan struct{}, 1),
		logger:      lol.NewZerolog(),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *WSClient) Dial(ctx context.Context) error {
	return c.connect(ctx)
}

func (c *WSClient) Close() error {
	var err error
	c.closeOnce.Do(func() {
		err = c.close()
	})
	return err
}

func (c *WSClient) close() error {
	close(c.done)

	err := c.dropActiveConnection()

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, s := range c.sharedSubscriptionByKey {
		s.clear()
	}

	return err
}
