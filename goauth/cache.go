package goauth

import (
	"time"
)

type StateCache interface {
	IsValid(string) bool
	Save(string, string, int64)
}

type State struct {
	ClientId  string
	ExpiredAt int64
}

type InMemoryStateCache struct {
	cachedStates map[string]State
	ticker       *time.Ticker
}

func NewInMemoryStateCache() *InMemoryStateCache {
	c := &InMemoryStateCache{
		cachedStates: make(map[string]State),
		ticker:       time.NewTicker(time.Duration(20)),
	}
	c.Start()
	return c
}

func (c *InMemoryStateCache) Start() {
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.timeoutChecker()
			}
		}
	}()
}

func (c *InMemoryStateCache) Stop() {
	c.ticker.Stop()
}

func (c *InMemoryStateCache) timeoutChecker() {
	now := time.Now().Unix()
	for k, v := range c.cachedStates {
		if v.ExpiredAt >= now {
			delete(c.cachedStates, k)
		}
	}
}

func (c *InMemoryStateCache) IsValid(state string) bool {
	_, ok := c.cachedStates[state]
	return ok
}

func (c *InMemoryStateCache) Save(state, clientId string, timeOut int64) {
	c.cachedStates[state] = State{
		ClientId:  clientId,
		ExpiredAt: time.Now().Unix() + timeOut,
	}
}
