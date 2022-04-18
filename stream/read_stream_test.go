package stream

import (
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestReadStream(t *testing.T) {
	t.Logf("\"ReadStream\" function test")

	const testFile = "./read_stream_sample.txt"
	const nGoroutines = 2
	const chunkSize = 16

	result := make([]byte, 0)
	expected, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	ch := make(chan []byte)
	wg := &sync.WaitGroup{}

	wg.Add(nGoroutines)

	go func() {
		defer wg.Done()
		defer close(ch)

		err = ReadStream(testFile, chunkSize, ch)
	}()

	go func() {
		defer wg.Done()

		i := 0
		for bytes := range ch {
			result = append(result, bytes...)
			i++
		}
	}()

	wg.Wait()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, result) {
		log.Fatalf(
			"File bytes not equal.\nExpected: %s\n     Got: %s",
			formatSlice(expected),
			formatSlice(result),
		)
	}
}

func TestReadStreamPartial(t *testing.T) {
	t.Logf("\"ReadStreamPartial\" function test")

	const testFile = "./read_stream_sample.txt"
	const nGoroutines = 2
	const chunkSize = 16
	const nChunks = 2
	const expectedLen = chunkSize * nChunks

	result := make([]byte, 0)
	expected, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}
	expected = expected[:expectedLen]

	ch := make(chan []byte)
	wg := &sync.WaitGroup{}

	wg.Add(nGoroutines)

	go func() {
		defer wg.Done()
		defer close(ch)

		err = ReadStreamPartial(testFile, nChunks, chunkSize, ch)
	}()

	go func() {
		i := 0
		for bytes := range ch {
			result = append(result, bytes...)
			i++
		}

		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, result) {
		log.Fatalf(
			"File bytes not equal.\nExpected: %s\n     Got: %s",
			formatSlice(expected),
			formatSlice(result),
		)
	}
}
