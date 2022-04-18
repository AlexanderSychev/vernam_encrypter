package transform

// Encrypt receive source bytes slice and key bytes slice and encrypts source slice
func Encrypt(src, key []byte) ([]byte, error) {
	srcLen, keyLen := len(src), len(key)

	if srcLen > keyLen {
		return make([]byte, 0), KeyToShortError{
			srcLen: srcLen,
			keyLen: keyLen,
		}
	}

	result := make([]byte, srcLen)

	for i := 0; i < srcLen; i++ {
		result[i] = src[i] ^ key[i]
	}

	return result, nil
}
