package internal

// Unique removes duplicate elements from a slice. Order is guaranteed to be preserved.
//
// Parameters:
//   - slice: The slice to remove duplicates from.
//
// Returns:
//   - []T: The filtered slice.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	for i := 0; i < len(slice)-1; i++ {
		elem := slice[i]

		top := i + 1

		for j := i + 1; j < len(slice); j++ {
			if slice[j] != elem {
				slice[top] = slice[j]
				top++
			}
		}

		slice = slice[:top:top]
	}

	return slice
}
