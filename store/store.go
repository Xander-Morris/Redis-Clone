package store

import (
	"fmt"
	"sync"
)

var (
	mu    sync.RWMutex
	Store = make(map[string]any)
)

func Set(key string, val any) {
	mu.Lock()
	defer mu.Unlock()
	Store[key] = val
}

// Used for either addition or subtraction by any provided number 
func Increment(key string, increment int) (int, error) {
	mu.Lock()
	defer mu.Unlock() 

	val, ok := Store[key]

	if !ok {
		return 0, fmt.Errorf("Key %s does not exist!", key)
	}

	actualVal, convOk := val.(int)

	if !convOk {
		return 0, fmt.Errorf("Key %s does not have an integer value!", key)
	}

	newVal := actualVal + increment
	Store[key] = newVal

	return newVal, nil 
}

func Delete(key string) {
	mu.Lock() 
	defer mu.Unlock() 
	delete(Store, key)
}

func Get(key string) (any, bool) {
	mu.RLock()
	defer mu.RUnlock()
	val, ok := Store[key]
	return val, ok
}
