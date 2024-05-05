package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Error("key not found")
		}
		if string(actual) != string(cas.inputVal) {
			t.Error("value doesn't match")
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
