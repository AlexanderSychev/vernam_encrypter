package stream

import (
	"bufio"
	"io"
	"os"
)

// cloneSlice is util function which makes clone of bytes slice - with same length and values
// but with other addresses in memory (need to avoid goroutines conflict)
func cloneSlice(src []byte) []byte {
	result := make([]byte, len(src))

	for i, b := range src {
		result[i] = b
	}

	return result
}

// ReadStream opens file, read it by chunks and send received bytes to channel
func ReadStream(filepath string, chunkSize int, ch chan<- []byte) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}

		ch <- cloneSlice(buf[:n])
	}
}

// ReadStreamPartial opens file, read received number of chunks and send received bytes to channel
func ReadStreamPartial(filepath string, nChunks, chunkSize int, ch chan<- []byte) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	for i := 0; i < nChunks; i++ {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		ch <- cloneSlice(buf[:n])
	}

	return nil
}
