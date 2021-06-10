package types

import "sync"

type ClientCache struct {
	clients map[string]*Client
	mutex   *sync.RWMutex
}

func NewClientCache() *ClientCache {
	return &ClientCache{
		clients: map[string]*Client{},
		mutex:   &sync.RWMutex{},
	}
}

func (c *ClientCache) Set(id string, client *Client) {
	c.mutex.Lock()
	c.clients[id] = client
	c.mutex.Unlock()
}

func (c *ClientCache) Get(id string) *Client {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.clients[id]
}

func (c *ClientCache) Delete(id string) {
	c.mutex.Lock()
	delete(c.clients, id)
	c.mutex.Unlock()
}

func (c *ClientCache) Run(do func(clients map[string]*Client)) {
	c.mutex.Lock()
	do(c.clients)
	c.mutex.Unlock()
}

func (c *ClientCache) Each(do func(id string, client *Client)) {
	c.mutex.RLock()
	for id, client := range c.clients {
		do(id, client)
	}
	c.mutex.RUnlock()
}
