package transform

import "fmt"

// KeyToShortError returns by Encrypt function when key length is lower than source length
type KeyToShortError struct {
	// Source bytes slice length
	srcLen int
	// Key bytes slice length
	keyLen int
}

func (e KeyToShortError) Error() string {
	return fmt.Sprintf("Key length is to short: expected at least %d bytes, got %d bytes", e.srcLen, e.keyLen)
}
