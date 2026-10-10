package store

import "sync"

var (
	mu    sync.RWMutex
	Store = make(map[string]any)
)

func Set(key string, val any) {
	mu.Lock()
	defer mu.Unlock()
	Store[key] = val
}

func Get(key string) (any, bool) {
	mu.RLock()
	defer mu.RUnlock()
	val, ok := Store[key]
	return val, ok
}
