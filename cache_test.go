package lrugo_test

import (
	"lrugo"
	"testing"

	"github.com/matryer/is"
)

func TestInsert(t *testing.T) {
	t.Run("Should insert an item", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](5)

		key := "uai"
		expected := "so"

		lru.Insert(key, expected)

		actual, _ := lru.Get(key)

		is.Equal(expected, actual)
	})

	t.Run("Should insert a second item", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](5)

		key := "uai"
		expected := "so"
		lru.Insert(key, expected)

		key2 := "trem"
		expected2 := "bao"
		lru.Insert(key2, expected2)

		actual, _ := lru.Get(key)
		actual2, _ := lru.Get(key2)

		is.Equal(expected, actual)
		is.Equal(expected2, actual2)
	})

	t.Run("Should insert a third item", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](5)

		key := "uai"
		expected := "so"
		lru.Insert(key, expected)

		key2 := "trem"
		expected2 := "bao"
		lru.Insert(key2, expected2)

		key3 := "tri"
		expected3 := "massa"
		lru.Insert(key3, expected3)

		actual, _ := lru.Get(key)
		actual2, _ := lru.Get(key2)
		actual3, _ := lru.Get(key3)

		is.Equal(expected, actual)
		is.Equal(expected2, actual2)
		is.Equal(expected3, actual3)
	})

	t.Run("Should remove the least used if full", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](3)

		input := []struct {
			key, value string
		}{
			{"uai", "so"},
			{"trem", "bao"},
			{"tri", "massa"},
		}

		for _, v := range input {
			lru.Insert(v.key, v.value)
		}

		key := "bah"
		value := "tche"
		lru.Insert(key, value)

		actual, ok := lru.Get(input[0].key)

		is.Equal("", actual)
		is.Equal(false, ok)
	})
}

func TestGet(t *testing.T) {
	t.Run("Should have a cache miss if key is not present", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](3)

		input := []struct {
			key, value string
		}{
			{"uai", "so"},
			{"trem", "bao"},
			{"tri", "massa"},
		}

		for _, v := range input {
			lru.Insert(v.key, v.value)
		}

		actual, ok := lru.Get("not here")

		is.Equal("", actual)
		is.Equal(false, ok)
	})

	t.Run("Should promote the used key when fetched", func(t *testing.T) {
		is := is.New(t)

		lru := lrugo.NewLRU[string](3)

		input := []struct {
			key, value string
		}{
			{"uai", "so"},
			{"trem", "bao"},
			{"tri", "massa"},
		}

		for _, v := range input {
			lru.Insert(v.key, v.value)
		}

		actual, ok := lru.Get(input[0].key)

		expected := input[0].value
		is.Equal(expected, actual)
		is.Equal(true, ok)

		key := "bah"
		value := "tche"
		lru.Insert(key, value)

		removed, ok := lru.Get(input[1].key)

		is.Equal("", removed)
		is.Equal(false, ok)
	})
}
