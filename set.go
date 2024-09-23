package sets

type Set interface {
	// IsEmpty checks whether the set is empty.
	//
	// Returns:
	//   - bool: True if the set is empty, false otherwise.
	IsEmpty() bool

	// Size returns the number of elements in the set.
	//
	// Returns:
	//   - int: The number of elements in the set. Never returns a negative number.
	Size() int

	// Reset removes all elements from the set.
	Reset()
}
