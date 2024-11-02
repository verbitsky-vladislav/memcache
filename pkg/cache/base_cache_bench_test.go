package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkBaseSet(b *testing.B) {
	cache := NewBaseInMemory[interface{}, interface{}]()

	numKeys := 100000
	keys := make([]string, numKeys)
	values := make([]interface{}, numKeys)
	for i := 0; i < numKeys; i++ {
		keys[i] = strconv.Itoa(rand.Intn(numKeys))
		values[i] = fmt.Sprintf("value-%d", i)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Выбираем случайный индекс для ключа и значения
			idx := rand.Intn(numKeys)
			cache.Set(keys[idx], values[idx])
		}
	})
}

func BenchmarkBaseGet(b *testing.B) {
	cache := NewBaseInMemory[interface{}, interface{}]()

	// Предварительно добавляем ключи и значения
	numKeys := 1000
	keys := make([]string, numKeys)
	for i := 0; i < numKeys; i++ {
		keys[i] = strconv.Itoa(i)
		cache.Set(keys[i], fmt.Sprintf("value-%s", keys[i]))
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Выбираем случайный ключ из предварительно созданных
			key := keys[rand.Intn(numKeys)]
			_, ok := cache.Get(key)
			if !ok {
				_ = fmt.Errorf("error getting value from cache by key: %s", key)
			}
		}
	})
}

func BenchmarkBaseSetGet(b *testing.B) {
	cache := NewBaseInMemory[interface{}, interface{}]()

	// Предварительная генерация ключей и значений
	numKeys := 100000
	keys := make([]string, numKeys)
	values := make([]interface{}, numKeys)
	for i := 0; i < numKeys; i++ {
		keys[i] = strconv.Itoa(i)
		values[i] = fmt.Sprintf("value-%d", i)
	}

	// Заполняем часть ключей, чтобы было что читать
	for i := 0; i < numKeys/2; i++ {
		cache.Set(keys[i], values[i])
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Выбираем случайный индекс
			idx := rand.Intn(numKeys)

			// 50% операций — запись, 50% — чтение
			if rand.Intn(2) == 0 {
				cache.Set(keys[idx], values[idx])
			} else {
				cache.Get(keys[idx])
			}
		}
	})
}
