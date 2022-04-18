package cmd

import (
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"sync"
	"vernam_encrypter/stream"
	"vernam_encrypter/transform"
)

const (
	// decryptSourceFileArgIndex is index of source file argument for "decrypt" command
	decryptSourceFileArgIndex = 0
	// decryptKeyFileArgIndex is index of key file argument for "decrypt" command
	decryptKeyFileArgIndex = 1
	// decryptTargetFileArgIndex is index of target file argument for "decrypt" command
	decryptTargetFileArgIndex = 2
	// decryptNArgs is number of arguments for "decrypt" command
	decryptNArgs = 3
	// decryptNGoroutines is number of parallel goroutines to run and wait by actionDecrypt function
	decryptNGoroutines = 4
)

// actionDecrypt is handler function for "decrypt" command
func actionDecrypt(c *cli.Context) error {
	if c.NArg() < decryptNArgs {
		return NotEnoughArgumentsError{
			command: "decrypt",
			requiredNArgs: decryptNArgs,
			receivedNArgs: c.NArg(),
		}
	}

	sourceFile, keyFile, targetFile :=
		c.Args().Get(decryptSourceFileArgIndex),
		c.Args().Get(decryptKeyFileArgIndex),
		c.Args().Get(decryptTargetFileArgIndex)

	sf, err := filepath.Abs(sourceFile)
	if err != nil {
		return err
	}

	kf, err := filepath.Abs(keyFile)
	if err != nil {
		return err
	}

	tf, err := filepath.Abs(targetFile)
	if err != nil {
		return err
	}

	sfStat, err := os.Stat(sf)
	if err != nil {
		return err
	}

	kfStat, err := os.Stat(kf)
	if err != nil {
		return err
	}

	if sfStat.Size() > kfStat.Size() {
		return KeyFileToShort{
			sourceFile: sf,
			keyFile: kf,
			sourceSize: sfStat.Size(),
			keySize: kfStat.Size(),
		}
	}

	nChunks, _ := calcChunks(int(sfStat.Size()))

	sourceChannel := make(chan []byte)
	keyChannel := make(chan []byte)
	decryptChannel := make(chan []byte)

	var wg sync.WaitGroup

	wg.Add(decryptNGoroutines)

	go func() {
		defer close(sourceChannel)
		defer wg.Done()

		err = stream.ReadStream(sf, chunkSize, sourceChannel)
	}()

	go func() {
		defer close(keyChannel)
		defer wg.Done()

		err = stream.ReadStreamPartial(kf, nChunks, chunkSize, keyChannel)
	}()

	go func() {
		defer close(decryptChannel)
		defer wg.Done()

		err = stream.TransformStream(sourceChannel, keyChannel, decryptChannel, transform.Decrypt)
	}()

	go func() {
		defer wg.Done()

		err = stream.WriteStream(tf, decryptChannel)
	}()

	wg.Wait()

	return nil
}
