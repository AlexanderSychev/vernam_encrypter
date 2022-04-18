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
	// encryptSourceFileArgIndex is index of source file argument for "encrypt" command
	encryptSourceFileArgIndex = 0
	// encryptKeyFileArgIndex is index of key file argument for "encrypt" command
	encryptKeyFileArgIndex = 1
	// encryptTargetFileArgIndex is index of target file argument for "encrypt" command
	encryptTargetFileArgIndex = 2
	// encryptNArgs is number of arguments for "encrypt" command
	encryptNArgs = 3
	// encryptNGoroutines is number of parallel goroutines to run and wait by actionEncrypt function
	encryptNGoroutines = 4
)

// actionEncrypt is handler function for "encrypt" command
func actionEncrypt(c *cli.Context) error {
	if c.NArg() < encryptNArgs {
		return NotEnoughArgumentsError{
			command: "encrypt",
			requiredNArgs: encryptNArgs,
			receivedNArgs: c.NArg(),
		}
	}

	sourceFile, keyFile, targetFile :=
		c.Args().Get(encryptSourceFileArgIndex),
		c.Args().Get(encryptKeyFileArgIndex),
		c.Args().Get(encryptTargetFileArgIndex)

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
	encryptChannel := make(chan []byte)

	var wg sync.WaitGroup

	wg.Add(encryptNGoroutines)

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
		defer close(encryptChannel)
		defer wg.Done()

		err = stream.TransformStream(sourceChannel, keyChannel, encryptChannel, transform.Encrypt)
	}()

	go func() {
		defer wg.Done()

		err = stream.WriteStream(tf, encryptChannel)
	}()

	wg.Wait()

	return nil
}
