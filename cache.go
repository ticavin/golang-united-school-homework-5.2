package cache

import "time"

type Istruct struct {
	value    string
	zeroHour time.Time
}

type Cache struct {
	data map[string]Istruct
}

func NewCache() Cache {
	return Cache{
		data: map[string]Istruct{}}
}

func (cache *Cache) Get(key string) (string, bool) {
	data, ok := cache.data[key]
	if !data.zeroHour.IsZero() && data.zeroHour.Before(time.Now()) {
		delete(cache.data, key)
		return "", false
	}
	return data.value, ok
}

func (cache *Cache) Put(key, value string) {
	cache.data[key] = Istruct{value: value}
}

func (cache *Cache) Keys() []string {
	keys := make([]string, 0, len(cache.data))
	for k := range cache.data {
		keys = append(keys, k)
	}
	return keys

}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.data[key] = Istruct{value: value, zeroHour: deadline}
}
