package main

import (
	"fmt"
	"math/rand"
	"memcache/pkg/cache"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// generateRandomString создает случайную строку из символов letters длиной от 5 до 10 символов
func generateRandomString(minLen, maxLen int) string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLen-minLen+1) + minLen
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	c := cache.NewShardedInMemory[interface{}, interface{}]()

	// Генерация 50-100 случайных ключей и значений
	numEntries := rand.Intn(51) + 50 // случайное число от 50 до 100
	for i := 0; i < numEntries; i++ {
		key := generateRandomString(5, 10)
		value := generateRandomString(5, 10)
		c.Set(key, value)
	}

	fmt.Println("Содержимое всех шардов:")
	c.GetAll()
}
