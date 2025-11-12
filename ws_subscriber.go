package hyperliquid

import "sync"

type callback func(any)

type sharedSubscription struct {
	id                  string
	count               int64
	subscriberByID      map[string]callback
	subscriberFunc      func(subscriptable)
	unsubscriberFunc    func(subscriptable)
	subscriptionPayload subscriptable

	mu sync.RWMutex
}

func newSubscriber(
	id string,
	payload subscriptable,
	subscriberFunc, unsubscriberFunc func(subscriptable),
) *sharedSubscription {
	return &sharedSubscription{
		id:                  id,
		subscriptionPayload: payload,
		count:               0,
		subscriberByID:      make(map[string]callback),
		subscriberFunc:      subscriberFunc,
		unsubscriberFunc:    unsubscriberFunc,
	}
}

func (u *sharedSubscription) addSubscriber(id string, cb callback) {
	u.mu.Lock()
	if _, exists := u.subscriberByID[id]; exists {
		u.mu.Unlock()
		return
	}
	u.subscriberByID[id] = cb
	u.count++
	c := u.count
	u.mu.Unlock()

	if c == 1 {
		u.subscriberFunc(u.subscriptionPayload)
	}
}

func (u *sharedSubscription) removeSubscriber(id string) {
	u.mu.Lock()
	if _, exists := u.subscriberByID[id]; !exists {
		u.mu.Unlock()
		return
	}
	delete(u.subscriberByID, id)
	c := u.count - 1
	u.count = c
	u.mu.Unlock()

	if c == 0 {
		u.unsubscriberFunc(u.subscriptionPayload)
	}
}

func (u *sharedSubscription) dispatch(data any) {
	u.mu.RLock()
	callbacks := make([]callback, 0, len(u.subscriberByID))
	for _, cb := range u.subscriberByID {
		callbacks = append(callbacks, cb)
	}
	u.mu.RUnlock()

	for _, cb := range callbacks {
		cb(data)
	}
}

func (u *sharedSubscription) clear() {
	u.mu.Lock()
	defer u.mu.Unlock()

	for id := range u.subscriberByID {
		delete(u.subscriberByID, id)
	}
	u.count = 0
	u.unsubscriberFunc(u.subscriptionPayload)
}

func (u *sharedSubscription) activePayload() (subscriptable, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	if u.count == 0 {
		return nil, false
	}

	return u.subscriptionPayload, true
}
