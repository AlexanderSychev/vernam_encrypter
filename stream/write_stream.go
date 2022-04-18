package stream

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func formatSlice(slice []byte) string {
	var builder strings.Builder

	builder.WriteString("[")
	for i, b := range slice {
		if i != 0 {
			builder.WriteString(", ")
		}
		formatted := strconv.FormatInt(int64(b), 16)
		if len([]rune(formatted)) < 2 {
			builder.WriteString("0")
		}

		builder.WriteString(formatted)
	}
	builder.WriteString("]")

	return builder.String()
}

// WriteStream read received chunks from channel and write them into file until channel will be closed
func WriteStream(filepath string, ch <-chan []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	writer := bufio.NewWriter(file)

	i := 0
	for buf := range ch {
		_, err := writer.Write(buf)

		if err != nil {
			return err
		}

		i++
	}

	return writer.Flush()
}
