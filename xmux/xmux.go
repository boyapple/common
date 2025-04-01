package xmux

import (
	"errors"
	"fmt"
	"sync"
)

// Mux 路由
type Mux[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

// New 创建一个mux实例
func New[K comparable, V any]() *Mux[K, V] {
	return &Mux[K, V]{
		m:    make(map[K]V),
		lock: sync.RWMutex{},
	}
}

// Register 注册,推荐在init函数使用
func (mux *Mux[K, V]) Register(key K, value V) {
	mux.lock.Lock()
	defer mux.lock.Unlock()
	_, ok := mux.m[key]
	if ok {
		panic(fmt.Sprintf("mux key[%v] already register", key))
	}
	mux.m[key] = value
}

// Get 获取路由
func (mux *Mux[K, V]) Get(key K) (V, error) {
	mux.lock.RLock()
	defer mux.lock.RUnlock()
	v, ok := mux.m[key]
	if ok {
		return v, nil
	}
	var empty V
	return empty, errors.New("key not found")
}
