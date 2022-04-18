package keygen

import "fmt"

// SliceToShortError returns by RandomBytes function when received length of bytes slice is two short
type SliceToShortError struct {
	// Received length of slice to generate
	length int
}

func (e SliceToShortError) Error() string {
	return fmt.Sprintf("Cannot generate bytes slice with length %d. Length must be at least 1 byte", e.length)
}
