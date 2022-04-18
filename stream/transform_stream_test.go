package stream

import (
	"reflect"
	"sync"
	"testing"
	"vernam_encrypter/transform"
)

func TestTransformStream(t *testing.T) {
	t.Logf("\"TransformStream\" function test")

	sourceStreamChunks := [][]byte{
		{15, 2, 55},
		{19, 7, 11},
	}
	keyStreamChunks := [][]byte{
		{10, 4, 22},
		{25, 34, 125},
	}
	expected := []byte{5, 6, 33, 10, 37, 118}
	result := make([]byte, 0)
	var err error = nil

	sourceCh, keyCh, transformCh := make(chan []byte), make(chan []byte), make(chan []byte)
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer close(sourceCh)
		defer wg.Done()

		for _, chunk := range sourceStreamChunks {
			sourceCh <- chunk
		}
	}()

	go func() {
		defer close(keyCh)
		defer wg.Done()

		for _, chunk := range keyStreamChunks {
			keyCh <- chunk
		}
	}()

	go func() {
		defer close(transformCh)
		defer wg.Done()

		err = TransformStream(sourceCh, keyCh, transformCh, transform.Encrypt)
	}()

	go func() {
		defer wg.Done()

		for ch := range transformCh {
			result = append(result, ch...)
		}
	}()

	wg.Wait()

	if err != nil {
		t.Fatalf("Got error: %s", err.Error())
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v got %v", expected, result)
	}
}
