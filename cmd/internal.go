package cmd

import (
	"math"
)

// chunkSize is minimal size of chunk in bytes
const chunkSize = 1024 * 1024

// calcChunks calculate number of bytes chunks for received file size.
func calcChunks(length int) (nChunks, lastChunkSize int) {
	nChunks = int(math.Ceil(float64(length) / float64(chunkSize)))
	lastChunkSize = nChunks * chunkSize - length

	return nChunks, lastChunkSize
}
