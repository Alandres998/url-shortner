package syncservices

import "sync"

type URLMap struct {
	sync.RWMutex
	m map[string]string
}

var URLStorage URLMap

func InitURLStorage() {
	URLStorage.m = make(map[string]string)
}

func (Store URLMap) Set(key string, value string) {
	Store.Lock()
	Store.m[key] = value
	Store.Unlock()
}

func (Store URLMap) Get(key string) (string, bool) {
	Store.RLock()
	originalURL, exists := Store.m[key]
	Store.RUnlock()
	return originalURL, exists
}
