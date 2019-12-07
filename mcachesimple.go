package mcachesimple

import (
	"errors"
	"sync"
)

type writeFuncType func(string, interface{})
type removeFuncType func(string)
type readFuncType func(string) (interface{}, error)

type item struct {
	key   string
	value interface{}
}

type Cache struct {
	items      map[string]*item
	writeFunc  writeFuncType
	removeFunc removeFuncType
	readFunc   readFuncType
	sync.Mutex
}

func New() *Cache {
	data := new(Cache)
	data.items = make(map[string]*item)
	return data
}

func (c *Cache) SetWriteFunc(fn writeFuncType) {
	c.writeFunc = fn
}

func (c *Cache) SetRemoveFunc(fn removeFuncType) {
	c.removeFunc = fn
}

func (c *Cache) SetReadFunc(fn readFuncType) {
	c.readFunc = fn
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()

	if data, found := c.items[key]; found {
		return data.value, nil

	} else {
		if c.readFunc != nil {
			readData, err := c.readFunc(key)

			if err != nil {
				return nil, err
			}

			c.items[key] = &item{key: key, value: readData}
			return readData, nil

		} else {
			return nil, errors.New("No data found error!")
		}
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	if data, found := c.items[key]; found {
		(*data).value = value

	} else {
		c.items[key] = &item{key: key, value: value}

	}

	if c.writeFunc != nil {
		c.writeFunc(key, value)

	}
}

func (c *Cache) Remove(key string) {
	c.Lock()
	defer c.Unlock()

	delete(c.items, key)

	if c.removeFunc != nil {
		c.removeFunc(key)

	}
}

func (c *Cache) Len() int {
	c.Lock()
	defer c.Unlock()

	return len(c.items)
}

func (c *Cache) Keys() []string {
	c.Lock()
	defer c.Unlock()

	keys := make([]string, len(c.items))
	i := 0
	for key, _ := range c.items {
		keys[i] = key
		i++
	}
	return keys
}
