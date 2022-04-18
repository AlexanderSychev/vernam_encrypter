package keygen

import "math/rand"

const maxRand = 256

// RandomBytes generates slice of random bytes
func RandomBytes(length int) ([]byte, error) {
	if length < 1 {
		return make([]byte, 0), SliceToShortError{
			length: length,
		}
	}

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = byte(rand.Intn(maxRand))
	}

	return result, nil
}
