package lrugo_test

import (
	"lrugo"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("Should insert an item", func(t *testing.T) {
		lru := lrugo.NewLRU()

		err := lru.Insert(key, value)
	})
}
