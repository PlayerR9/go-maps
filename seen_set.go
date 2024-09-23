package sets

import (
	"github.com/PlayerR9/go-sets/internal"
)

// SeenSet is a map that keeps track of seen values.
type SeenSet[T comparable] struct {
	// table is the map that keeps track of seen values.
	table map[T]struct{}
}

// IsEmpty implements the Set interface.
func (sm SeenSet[T]) IsEmpty() bool {
	return len(sm.table) == 0
}

// Size implements the Set interface.
func (sm SeenSet[T]) Size() int {
	return len(sm.table)
}

// Reset implements the Set interface.
func (sm *SeenSet[T]) Reset() {
	if sm == nil {
		return
	}

	if len(sm.table) > 0 {
		for k := range sm.table {
			delete(sm.table, k)
		}
	}
}

// NewSeenSet creates a new SeenSet.
//
// Returns:
//   - *SeenSet[T]: The new SeenSet. Never returns nil.
func NewSeenSet[T comparable]() *SeenSet[T] {
	return &SeenSet[T]{
		table: make(map[T]struct{}),
	}
}

// See sets the value as seen.
//
// Parameters:
//   - v: The value to set as seen.
//
// Returns:
//   - bool: True if the value was set as seen or the receiver was nil. False otherwise.
func (sm *SeenSet[T]) See(v T) bool {
	if sm == nil {
		return false
	}

	_, ok := sm.table[v]
	if ok {
		return false
	}

	sm.table[v] = struct{}{}

	return true
}

// SetSeen sets the value as seen. Does nothing if the receiver is nil
// or the value is already seen.
//
// Parameters:
//   - v: The value to set as seen.
func (sm *SeenSet[T]) SetSeen(v T) {
	if sm == nil {
		return
	}

	_, ok := sm.table[v]
	if ok {
		return
	}

	sm.table[v] = struct{}{}
}

// Has checks whether the value is seen.
//
// Parameters:
//   - v: The value to check.
//
// Returns:
//   - bool: True if the value is seen, false otherwise.
func (sm SeenSet[T]) Has(v T) bool {
	_, ok := sm.table[v]
	return ok
}

// FilterSeen returns the elements that are seen. The order of the elements is preserved
// and no duplicates are contained.
//
// Parameters:
//   - elems: The elements to filter.
//
// Returns:
//   - []T: The elements that are seen.
func (sm SeenSet[T]) FilterSeen(elems []T) []T {
	slice := make([]T, 0, len(elems))

	for i := 0; i < len(elems); i++ {
		_, ok := sm.table[elems[i]]
		if ok {
			slice = append(slice, elems[i])
		}
	}

	slice = internal.Unique(slice)
	return slice
}

// FilterNotSeen is like FilterSeen but returns the elements that are not seen.
//
// Parameters:
//   - elems: The elements to filter.
//
// Returns:
//   - []T: The elements that are not seen.
func (sm SeenSet[T]) FilterNotSeen(elems []T) []T {
	slice := make([]T, 0, len(elems))

	for i := 0; i < len(elems); i++ {
		_, ok := sm.table[elems[i]]
		if !ok {
			slice = append(slice, elems[i])
		}
	}

	slice = internal.Unique(slice)
	return slice
}
