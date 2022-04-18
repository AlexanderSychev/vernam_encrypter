package transform

// Decrypt receive source bytes slice and key bytes slice and decrypts source slice
func Decrypt(src, key []byte) ([]byte, error) {
	srcLen, keyLen := len(src), len(key)

	if srcLen > keyLen {
		return make([]byte, 0), KeyToShortError{
			srcLen: srcLen,
			keyLen: keyLen,
		}
	}

	result := make([]byte, srcLen)

	for i := 0; i < srcLen; i++ {
		result[i] = key[i] ^ src[i]
	}

	return result, nil
}
