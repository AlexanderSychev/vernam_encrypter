package cmd

import (
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"vernam_encrypter/keygen"
	"vernam_encrypter/stream"
)

// keygenNArgs is required number of "keygen" command arguments
const keygenNArgs = 2

// keygenNGoroutines is number of parallel key generation goroutines to wait
const keygenNGoroutines = 2

// actionKeygen is handler function for "keygen" command
func actionKeygen(c *cli.Context) error {
	if c.NArg() < keygenNArgs {
		return NotEnoughArgumentsError{
			command: "keygen",
			requiredNArgs: keygenNArgs,
			receivedNArgs: c.NArg(),
		}
	}

	sLength, filename := c.Args().Get(0), c.Args().Get(1)

	length, err := strconv.Atoi(sLength)
	if err != nil {
		return err
	}

	fp, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	if length > chunkSize {
		// If received length is bigger then chunk size (1 MB) then generate key file and write it by chunks

		nChunks, lastChunkSize := calcChunks(length)

		ch := make(chan []byte)
		var wg sync.WaitGroup
		wg.Add(keygenNGoroutines)

		// First goroutine will generate key bytes by chunks
		go func() {
			defer close(ch) // Close channel after all chunks generation (need to avoid deadlock)
			defer wg.Done()

			err = stream.RandomBytesStream(nChunks, chunkSize, lastChunkSize, ch)
		}()

		// Second goroutine will write received generated chunks to file
		go func() {
			defer wg.Done()

			err = stream.WriteStream(fp, ch)
		}()

		// We will wait until both goroutines will be finished
		wg.Wait()

		if err != nil {
			return err
		}
	} else {
		// Otherwise, generate key file immediately

		var bytes []byte
		bytes, err = keygen.RandomBytes(length)
		if err != nil {
			return err
		}

		err = os.WriteFile(fp, bytes, 0666)
		if err != nil {
			return err
		}
	}

	return nil
}
