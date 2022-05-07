package main

import (
	"errors"
	"fmt"
	"mp4Service/mp4"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
)

func process(data *Data) ([]byte, error) {
	switch data.Type {
	case processRequest:
		inputPath := string(data.Payload)

		// Read the input file
		reader, err := mp4.New(inputPath)
		if err != nil {
			return nil, err
		}

		// Get the init segment from file
		result, err := reader.GetInitSegment()
		if err != nil {
			return nil, err
		}

		// Create result file
		resultPath := _ResultPath
		resultFile, err := os.Create(resultPath)
		if err != nil {
			return nil, err
		}

		// Close result file on exit and check for its returned error
		defer func() {
			if err := resultFile.Close(); err != nil {
				panic(err)
			}
		}()

		// Write bytes to result file
		bytesWritten, err := resultFile.Write(result)
		if err != nil {
			return nil, err
		}

		// Return absolute result path
		absoluteResultPath, err := filepath.Abs(resultPath)
		if err != nil {
			return nil, err
		}
		fmt.Println(aurora.Sprintf(aurora.Green("Wrote %d bytes to %s"), bytesWritten, absoluteResultPath))

		return []byte(absoluteResultPath), nil
	default:
		return nil, errors.New("undefined type for data")
	}
}
