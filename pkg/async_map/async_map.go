package asyncmap

import (
	"sync"

	"github.com/vladjong/user_grade_api/pkg/hellpers"
)

type AsyncMap interface {
	Set(key string, value interface{}) error
	Get(key string) (value interface{}, err error)
	GetAll() (value []interface{}, err error)
	Delete(key string) error
}

type mutexMap struct {
	mx      sync.RWMutex
	storage map[string]interface{}
}

func NewCache() *mutexMap {
	return &mutexMap{
		storage: make(map[string]interface{}),
	}
}

func (c *mutexMap) Set(key string, value interface{}) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	return nil
}

func (c *mutexMap) Get(key string) (value interface{}, err error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.storage[key]
	if !ok {
		return "", hellpers.ErrNotFound
	}
	return value, nil
}

func (c *mutexMap) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}

func (c *mutexMap) GetAll() (values []interface{}, err error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, value := range c.storage {
		values = append(values, value)
	}
	return values, nil
}
