package cmd

import "fmt"

// NotEnoughArgumentsError returns by commands action functions when not enough arguments received
type NotEnoughArgumentsError struct {
	// Name of failed command
	command string
	// Required number of arguments
	requiredNArgs int
	// Number of received arguments
	receivedNArgs int
}

func (e NotEnoughArgumentsError) Error() string {
	return fmt.Sprintf(
		"Not enough arguments for command \"%s\": expected %d arguments, got %d",
		e.command,
		e.requiredNArgs,
		e.receivedNArgs,
	)
}

// KeyFileToShort returns by command action functions when key file size is lower than source file
type KeyFileToShort struct {
	// Absolute path to source file (which must be encrypted or decrypted)
	sourceFile string
	// Size of source file in bytes
	sourceSize int64
	// Absolute path ot key file (which must be used to encrypt or decrypt source file)
	keyFile string
	// Size of key file in bytes
	keySize int64
}

func (e KeyFileToShort) Error() string {
	return fmt.Sprintf(
		"Key file \"%s\" is to short for file \"%s\": expected at least %d bytes, got %d bytes",
		e.keyFile,
		e.sourceFile,
		e.sourceSize,
		e.keySize,
	)
}
