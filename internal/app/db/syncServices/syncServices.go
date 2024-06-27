package syncservices

import "sync"

type UrlMap struct {
	sync.RWMutex
	m map[string]string
}

var UrlStorage UrlMap

func InitUrlStorage() {
	UrlStorage.m = make(map[string]string)
}

func (Store UrlMap) Set(key string, value string) {
	Store.Lock()
	Store.m[key] = value
	Store.Unlock()
}

func (Store UrlMap) Get(key string) (string, bool) {
	Store.RLock()
	originalURL, exists := Store.m[key]
	Store.RUnlock()
	return originalURL, exists
}
