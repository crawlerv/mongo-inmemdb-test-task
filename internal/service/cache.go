package service

import (
	"sync"
)

var m sync.Map

type MemoryCache struct{}

func NewMemoryCache() MemoryCache {
	return MemoryCache{}
}

func (mc MemoryCache) Set(key interface{}, value interface{}) error {
	m.Store(key, value)

	return nil
}

func (mc MemoryCache) Update(key interface{}, value interface{}) error {
	mc.Delete(key)
	return mc.Set(key, value)
}

func (mc MemoryCache) Get(key interface{}) interface{} {
	r, _ := m.Load(key)
	return r
}

func (mc MemoryCache) Delete(key interface{}) {
	m.Delete(key)
}

func (mc MemoryCache) GetAll() []interface{} {
	var res []interface{}
	m.Range(func(key interface{}, value interface{}) bool {
		res = append(res, value)
		return true
	})
	return res
}

func (mc MemoryCache) IsExists(key interface{}) bool {
	_, r := m.Load(key)
	return r
}
