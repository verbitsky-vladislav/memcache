package cache

import "sync"

// BaseCache represents a cache interface,
// providing methods for working with data: getting, setting, and deleting values.
type BaseCache[K comparable, V any] interface {
	Get(k K) (V, bool) // get item from storage
	Set(k K, v V)      // set item to storage
	Delete(k K)        // delete item by key
	Load() int         // get storage length
	Clear()            // full clear storage
	GetAll() map[K]V   // return all items from storage
}

// BaseInMemory implements an in-memory cache,
// using mutexes to ensure safe access in a multithreaded environment.
type BaseInMemory[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewBaseInMemory creates a new instance of BaseInMemory,
// initializing the internal data storage.
func NewBaseInMemory[K comparable, V any]() *BaseInMemory[K, V] {
	return &BaseInMemory[K, V]{
		data: make(map[K]V),
	}
}

// Get retrieves the value associated with the key k.
// It returns the zero value of type V and false if the key is not found.
func (c *BaseInMemory[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.data[k]

	return v, ok
}

// Set stores the value v under the key k, ensuring write lock.
func (c *BaseInMemory[K, V]) Set(k K, v V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

// Delete removes the value associated with the key k.
func (c *BaseInMemory[K, V]) Delete(k K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, k)
}

// Load return num of elements from map
func (c *BaseInMemory[K, V]) Load() int {
	return len(c.data)
}

// Clear delete all data from map
func (c *BaseInMemory[K, V]) Clear() {
	for k := range c.data {
		delete(c.data, k)
	}
}

// GetAll return all items from map
func (c *BaseInMemory[K, V]) GetAll() map[K]V {
	allData := make(map[K]V)
	for k, v := range c.data {
		allData[k] = v
	}
	return allData
}
