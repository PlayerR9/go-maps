package sets

import (
	"cmp"
	"iter"
	"slices"
)

// OrderedMap is a map that is ordered by the keys.
type OrderedMap[K cmp.Ordered, V any] struct {
	// values is a map of the values in the map.
	values map[K]V

	// keys is a slice of the keys in the map.
	keys []K
}

// IsEmpty implements the Set interface.
func (m OrderedMap[K, V]) IsEmpty() bool {
	return len(m.keys) == 0
}

// Size implements the Set interface.
func (m OrderedMap[K, V]) Size() int {
	return len(m.keys)
}

// Reset implements the Set interface.
func (m *OrderedMap[K, V]) Reset() {
	if m == nil {
		return
	}

	if len(m.keys) == 0 {
		return
	}

	m.keys = m.keys[:0]

	for key := range m.values {
		delete(m.values, key)
	}
}

// NewOrderedMap creates a new OrderedMap.
//
// Returns:
//   - *OrderedMap: A pointer to the newly created OrderedMap. Never returns nil.
func NewOrderedMap[K cmp.Ordered, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		values: make(map[K]V),
		keys:   make([]K, 0),
	}
}

// Get returns the value of the key in the map.
//
// Parameters:
//   - key: The key to get.
//
// Returns:
//   - V: The value of the key in the map.
//   - bool: True if the key exists in the map. False if the key does not exist.
func (m OrderedMap[K, V]) Get(key K) (V, bool) {
	if len(m.keys) == 0 {
		return *new(V), false
	}

	value, ok := m.values[key]
	return value, ok
}

// Contains checks if the key exists in the map.
//
// Parameters:
//   - key: The key to check.
//
// Returns:
//   - bool: True if the key exists in the map. False if the key does not exist.
func (m OrderedMap[K, V]) Contains(key K) bool {
	if len(m.keys) == 0 {
		return false
	}

	_, ok := slices.BinarySearch(m.keys, key)
	return ok
}

// Remove removes the key from the map. Does nothing if the receiver is nil
// or the key does not exist.
//
// Parameters:
//   - key: The key to remove.
func (m *OrderedMap[K, V]) Remove(key K) {
	if m == nil {
		return
	}

	pos, ok := slices.BinarySearch(m.keys, key)
	if !ok {
		return
	}

	m.keys = slices.Delete(m.keys, pos, pos+1)
	delete(m.values, key)
}

// Add adds a key-value pair to the map. Does nothing if the receiver is nil
// or the key already exists.
//
// Parameters:
//   - key: The key to add.
//   - value: The value to add.
func (m *OrderedMap[K, V]) Add(key K, value V) {
	if m == nil {
		return
	}

	pos, ok := slices.BinarySearch(m.keys, key)
	if ok {
		return
	}

	m.keys = slices.Insert(m.keys, pos, key)
	m.values[key] = value
}

// ForceAdd is the same as Add, except that it will overwrite the value if
// the key already exists.
//
// Parameters:
//   - key: The key to add.
//   - value: The value to add.
func (m *OrderedMap[K, V]) ForceAdd(key K, value V) {
	if m == nil {
		return
	}

	pos, ok := slices.BinarySearch(m.keys, key)
	if !ok {
		m.keys = slices.Insert(m.keys, pos, key)
	}

	m.values[key] = value
}

// Map returns a copy of the map of the values in the map.
//
// Returns:
//   - map[K]V: The map of the values in the map. Nil if there are no keys.
func (m OrderedMap[K, V]) Map() map[K]V {
	if len(m.keys) == 0 {
		return nil
	}

	map_copy := make(map[K]V, len(m.values))

	for key, value := range m.values {
		map_copy[key] = value
	}

	return map_copy
}

// Keys returns a copy of the keys in the map.
//
// Returns:
//   - []K: The keys in the map.
func (m OrderedMap[K, V]) Keys() []K {
	if len(m.keys) == 0 {
		return nil
	}

	keys := make([]K, 0, len(m.keys))
	copy(keys, m.keys)

	return keys
}

// Entry returns an iterator that iterates over the entries in the map according
// to the order of the keys.
//
// Returns:
//   - iter.Seq2[K, V]: The iterator. Never returns nil.
func (m OrderedMap[K, V]) Entry() iter.Seq2[K, V] {
	if len(m.keys) == 0 {
		return func(yield func(K, V) bool) {}
	}

	return func(yield func(key K, value V) bool) {
		for _, key := range m.keys {
			if !yield(key, m.values[key]) {
				break
			}
		}
	}
}
