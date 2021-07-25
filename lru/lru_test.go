package lru

import (
	"math/rand"
	"testing"
)

func TestLRUCache(t *testing.T) {
	keyMax := int32(10)
	cache := Constructor(10)
	for i := int32(1); i <= keyMax; i++ {
		cache.Put(int(i), int(i)+1000)
	}

	for i := 0; i < 1000; i++ {
		cache.Get(int(rand.Int31n(keyMax)) + 1)
	}

	t.Logf("cache: %+v", *cache)
}
