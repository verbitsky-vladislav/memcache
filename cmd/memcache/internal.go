package main

type Cache[K comparable, V interface{}] interface {
	Get(k K) (v V, ok bool)
	Set(k K, v V)
	Delete(k K)
}

type CacheByString[K comparable, V interface{}] struct {
	store map[K]V
}

func (cache *CacheByString[K, V]) Get(k K) (value V, ok bool) {
	value, ok = cache.store[k]
	return
}

func (cache *CacheByString[K, V]) Set(k K, v V) {
	cache.store[k] = v
}

func (cache *CacheByString[K, V]) Delete(k K) {
	delete(cache.store, k)
}

func CreateMyCache[K comparable, V interface{}](cacheFunc func(K) string) *CacheByString[K, V] {
	store := make(map[K]V)
	var cache = CacheByString[K, V]{
		store: store,
	}
	return &cache
}
