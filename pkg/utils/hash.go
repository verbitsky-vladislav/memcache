package utils

import (
	"fmt"
)

func Hash[K comparable](key K, seed int32) int32 {
	//h := fnv.New32a()
	//h.Write([]byte(fmt.Sprintf("%v", key)))
	//return int32(h.Sum32()) ^ seed
	hash := seed
	for _, b := range []byte(fmt.Sprintf("%v", key)) {
		hash += int32(b)
	}
	return hash
}
