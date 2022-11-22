package asyncmap

import (
	"sync"

	"github.com/vladjong/user_grade_api/internal/entity"
	"github.com/vladjong/user_grade_api/pkg/hellpers"
)

type AsyncMap interface {
	Set(key string, value entity.UserGrade) error
	Get(key string) (entity.UserGrade, error)
	Delete(key string) error
}

type mutexMap struct {
	mx      sync.RWMutex
	storage map[string]entity.UserGrade
}

func NewCache() *mutexMap {
	return &mutexMap{
		storage: make(map[string]entity.UserGrade),
	}
}

func (c *mutexMap) Set(key string, value entity.UserGrade) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	return nil
}

func (c *mutexMap) Get(key string) (entity.UserGrade, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.storage[key]
	if !ok {
		return value, hellpers.ErrNotFound
	}
	return value, nil
}

func (c *mutexMap) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}
