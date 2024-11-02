package cache

import (
	"math/rand"
	"memcache/pkg/utils"
)

type ShardedCache[K comparable, V any] interface {
	Get(k K) (V, bool)
	Set(k K, v V)
	Delete(k K)
	GetAll() []*BaseInMemory[K, V]
}

type ShardedInMemory[K comparable, V any] struct {
	shards []*BaseInMemory[K, V]
	hash   func(K, int32) int32

	hash0     int32 // seed to avoid collisions
	numShards int32 // num of shards

}

func NewShardedInMemory[K comparable, V any]() *ShardedInMemory[K, V] {
	shards := make([]*BaseInMemory[K, V], 16)
	for i := 0; i < 16; i++ {
		shards[i] = NewBaseInMemory[K, V]()
	}

	return &ShardedInMemory[K, V]{
		shards:    shards,
		hash0:     rand.Int31(),
		hash:      utils.Hash[K],
		numShards: int32(len(shards)),
	}
}

func (s *ShardedInMemory[K, V]) Get(k K) (V, bool) {
	index := s.getShardIndex(k)
	val, ok := s.shards[index].Get(k)
	return val, ok
}

func (s *ShardedInMemory[K, V]) Set(k K, v V) {

	index := s.getShardIndex(k)

	s.shards[index].Set(k, v)
}

func (s *ShardedInMemory[K, V]) Delete(k K) {

	index := s.getShardIndex(k)
	s.shards[index].Delete(k)
}

func (s *ShardedInMemory[K, V]) getShardIndex(k K) int32 {
	hashValue := s.hash(k, s.hash0)
	return hashValue % s.numShards
}
