package goauth

import "time"

type StateCache interface {
	Get(string) string
	Set(string, string, int64)
}

type State struct {
	Value     string
	ExpiredAt int64
}

type InMemoryStateCache struct {
	cache map[string]State
}

func (c *InMemoryStateCache) Get(clientId string) string {
	time.Now().Unix()

	return c.cache[clientId].Value
}

func (c *InMemoryStateCache) Set(clientId, state string, timeOut int64) {
	c.cache[clientId] = State{
		Value:     state,
		ExpiredAt: time.Now().Unix() + timeOut,
	}
}
