package stream

import "vernam_encrypter/keygen"

// RandomBytesStream generates N chunks of random bytes and sends them to received channel
func RandomBytesStream(nChunks, chunkSize, lastChunkSize int, ch chan<- []byte) error {
	var bytes []byte
	var err error

	for i := 0; i < nChunks; i++ {
		randomLength := chunkSize
		if i == nChunks - 1 && lastChunkSize > 0 {
			randomLength = lastChunkSize
		}

		bytes, err = keygen.RandomBytes(randomLength)
		if err != nil {
			return err
		}

		ch <- bytes
	}

	return nil
}
