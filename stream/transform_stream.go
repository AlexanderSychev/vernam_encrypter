package stream

func TransformStream(sourceCh, keyCh, transformCh chan []byte, transformer func([]byte, []byte) ([]byte, error)) error {
	var err error = nil

	for {
		sourceChunk, sourceOpened := <- sourceCh
		keyChunk, keyOpened := <- keyCh

		if !sourceOpened && !keyOpened {
			break
		}

		var transformedChunk []byte
		transformedChunk, err = transformer(sourceChunk, keyChunk)

		transformCh <- transformedChunk
	}

	return err
}
