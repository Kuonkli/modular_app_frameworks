package core

import "strings"

type factoryEntry struct {
	kind    string
	factory func(*Container) interface{}
}

type Container struct {
	factories  map[string]factoryEntry
	singletons map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		factories:  make(map[string]factoryEntry),
		singletons: make(map[string]interface{}),
	}
}

func (c *Container) AddSingleton(key string, factory func(*Container) interface{}) {
	c.factories[key] = factoryEntry{"singleton", factory}
}

func (c *Container) AddTransient(key string, factory func(*Container) interface{}) {
	c.factories[key] = factoryEntry{"transient", factory}
}

func (c *Container) Get(key string) interface{} {
	entry, ok := c.factories[key]
	if !ok {
		panic("service not found: " + key)
	}

	if entry.kind == "singleton" {
		if val, ok := c.singletons[key]; ok {
			return val
		}
		instance := entry.factory(c)
		c.singletons[key] = instance
		return instance
	}

	return entry.factory(c)
}

func (c *Container) GetMany(prefix string) []interface{} {
	var res []interface{}
	for key := range c.factories {
		if strings.HasPrefix(key, prefix) {
			res = append(res, c.Get(key))
		}
	}
	return res
}
