package model

import (
	"encoding/json"
	"hash/fnv"
	"sync"
)

//ShardCount is the number of map shards
var ShardCount = 128

// TODO: Add Keys function which returns an array of keys for the map.

// ConcurrentMap is a "thread" safe map of type string:Service.
// To avoid lock bottlenecks this map is dived to several (ShardCount) map shards.
type ConcurrentMap []*ConcurrentMapShared

// ConcurrentMapShared is the struct used in ConcurrentMap
type ConcurrentMapShared struct {
	items        map[string]Service
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// NewCMap creates a new concurrent map.
func NewCMap() ConcurrentMap {
	m := make(ConcurrentMap, ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]Service)}
	}
	return m
}

//GetShard returns a shard based on the given key
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	hasher := fnv.New32()
	hasher.Write([]byte(key))
	return m[int(hasher.Sum32())%ShardCount]
}

// Set sets the given value to the specified key.
func (m *ConcurrentMap) Set(key string, value Service) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.items[key] = value
}

// Get retrieves an element from the map based on the given key.
func (m ConcurrentMap) Get(key string) (Service, bool) {
	// Get shard.
	shard := m.GetShard(key)
	shard.RLock()
	defer shard.RUnlock()

	// Get item from shard.
	val, ok := shard.items[key]
	return val, ok
}

// Count returns the number of elements within the map.
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < ShardCount; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has looks up an item under the specified key
func (m *ConcurrentMap) Has(key string) bool {
	// Get shard.
	shard := m.GetShard(key)
	shard.RLock()
	defer shard.RUnlock()

	// See if element is within shard.
	_, ok := shard.items[key]
	return ok
}

// Remove removes an element from the map based on the provided key.
func (m *ConcurrentMap) Remove(key string) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.items, key)
}

// IsEmpty checks if the map is empty.
func (m *ConcurrentMap) IsEmpty() bool {
	return m.Count() == 0
}

// Tuple is used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type Tuple struct {
	Key string
	Val Service
}

// Iter returns an iterator which could be used in a for range loop.
func (m ConcurrentMap) Iter() <-chan Tuple {
	ch := make(chan Tuple)

	go func() {
		// Iterate each shard.
		for _, shard := range m {
			shard.RLock()

			// Iterate each key, value pair.
			for key, val := range shard.items {
				ch <- Tuple{key, val}
			}

			shard.RUnlock()
		}
		close(ch)
	}()

	return ch
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m ConcurrentMap) IterBuffered() <-chan Tuple {
	ch := make(chan Tuple, m.Count())

	go func() {
		// Iterate each shard.
		for _, shard := range m {
			// Iterate each key, value pair.
			shard.RLock()

			for key, val := range shard.items {
				ch <- Tuple{key, val}
			}

			shard.RUnlock()
		}
		close(ch)
	}()

	return ch
}

// MarshalJSON reviles ConcurrentMap "private" variables to JSON marshal.
func (m ConcurrentMap) MarshalJSON() ([]byte, error) {
	// Create a temporary map, which will hold all item spread across shards.
	tmp := make(map[string]Service)

	// Insert items to temporary map.
	for item := range m.Iter() {
		tmp[item.Key] = item.Val
	}

	return json.Marshal(tmp)
}

//UnmarshalJSON is the reverse process of Marshal.
func (m *ConcurrentMap) UnmarshalJSON(b []byte) (err error) {
	tmp := make(map[string]Service)

	// Unmarshal into a single map.
	if err := json.Unmarshal(b, &tmp); err != nil {
		return nil
	}

	// foreach key,value pair in temporary map insert into our concurrent map.
	for key, val := range tmp {
		m.Set(key, val)
	}

	return nil
}

/*
CMAP license:

The MIT License (MIT)

Copyright (c) 2014 streamrail
https://github.com/streamrail/concurrent-map/
*/
