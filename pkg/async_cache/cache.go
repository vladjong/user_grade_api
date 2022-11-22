package asynccache

import (
	"sync"

	"github.com/vladjong/user_grade_api/internal/entity"
	"github.com/vladjong/user_grade_api/pkg/hellpers"
)

type Cache interface {
	Set(key string, value entity.UserGrade) error
	Get(key string) (entity.UserGrade, error)
	Delete(key string) error
}

type mutexCache struct {
	mx      sync.RWMutex
	storage map[string]entity.UserGrade
}

func NewCache() *mutexCache {
	return &mutexCache{
		storage: make(map[string]entity.UserGrade),
	}
}

func (c *mutexCache) Set(key string, value entity.UserGrade) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	return nil
}

func (c *mutexCache) Get(key string) (entity.UserGrade, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.storage[key]
	if !ok {
		return value, hellpers.ErrNotFound
	}
	return value, nil
}

func (c *mutexCache) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}
