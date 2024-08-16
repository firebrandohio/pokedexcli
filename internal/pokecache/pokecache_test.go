package pokecache

import (
	"testing"
	"time"
)

func TestCreateAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	if cache.cache == nil {
		t.Error("cache is nil")
		return
	}

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {

		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesnt match %s", string(actual), string(cas.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Error("value was not reaped")
	}
}
