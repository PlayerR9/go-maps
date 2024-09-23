package sets

import (
	"iter"
	"slices"
)

// EqualSet represents a set of elements that implements the Equals method.
//
// Interface: Equals(other T) bool
type EqualSet[T interface {
	Equals(other T) bool
}] struct {
	// elems is the set of elements
	elems []T
}

// IsEmpty implements the Set interface.
func (s EqualSet[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Size implements the Set interface.
func (s EqualSet[T]) Size() int {
	return len(s.elems)
}

// Reset implements the Set interface.
func (s *EqualSet[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.elems) > 0 {
		for i := 0; i < len(s.elems); i++ {
			s.elems[i] = *new(T)
		}
		s.elems = s.elems[:0]
	}
}

// NewEqualSet creates a new empty set.
//
// Returns:
//   - *EqualSet[T]: The created set. Never returns nil.
func NewEqualSet[T interface {
	Equals(other T) bool
}]() *EqualSet[T] {
	return &EqualSet[T]{
		elems: make([]T, 0),
	}
}

// Add adds an element to the set. If the element already exists or the receiver
// is nil, this function does nothing.
//
// Parameters:
//   - elem: The element to add.
func (s *EqualSet[T]) Add(elem T) {
	if s == nil {
		return
	}

	has_element := slices.ContainsFunc(s.elems, elem.Equals)
	if !has_element {
		s.elems = append(s.elems, elem)
	}
}

// AddMany is the same as Add, but adds many elements at once. More optimized
// than calling Add multiple times.
//
// Parameters:
//   - elems: The elements to add.
func (s *EqualSet[T]) AddMany(elems []T) {
	if s == nil || len(elems) == 0 {
		return
	}

	for i := 0; i < len(elems); i++ {
		has_element := slices.ContainsFunc(s.elems, elems[i].Equals)
		if !has_element {
			s.elems = append(s.elems, elems[i])
		}
	}
}

// Union adds all elements from another set to the set.
//
// Parameters:
//   - other: The other set to add.
//
// Returns:
//   - int: The number of elements added.
//
// If the receiver or 'other' is nil, then 0 is returned, always.
func (s *EqualSet[T]) Union(other *EqualSet[T]) int {
	if s == nil || other == nil {
		return 0
	}

	var count int

	for _, elem := range other.elems {
		if !slices.ContainsFunc(s.elems, elem.Equals) {
			s.elems = append(s.elems, elem)
			count++
		}
	}

	return count
}

// All returns an iterator that iterates over all elements in the set.
//
// Returns:
//   - iter.Seq[T]: The iterator. Never returns nil.
func (s *EqualSet[T]) All() iter.Seq[T] {
	var fn func(yield func(T) bool)

	if s == nil {
		fn = func(yield func(T) bool) {}
	} else {
		fn = func(yield func(T) bool) {
			for _, elem := range s.elems {
				if !yield(elem) {
					return
				}
			}
		}
	}

	return fn
}
