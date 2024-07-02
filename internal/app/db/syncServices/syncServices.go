package syncservices

import "sync"

type URLMap struct {
	s sync.RWMutex
	m map[string]string
}

var URLStorage URLMap

func InitURLStorage() {
	URLStorage.m = make(map[string]string)
}

func (Store *URLMap) Set(key string, value string) {
	Store.s.Lock()
	Store.m[key] = value
	Store.s.Unlock()
}

func (Store *URLMap) Get(key string) (string, bool) {
	Store.s.RLock()
	originalURL, exists := Store.m[key]
	Store.s.RUnlock()
	return originalURL, exists
}
