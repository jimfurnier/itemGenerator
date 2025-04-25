package compressor

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

type GzipCompressor struct{}

func (GzipCompressor) Compress(inputPath string) (string, error) {
	outputPath := inputPath + ".gz"

	inFile, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			panic(err)
		}
	}(inFile)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(err)
		}
	}(outFile)

	w := gzip.NewWriter(outFile)
	defer func(w *gzip.Writer) {
		err := w.Close()
		if err != nil {
			panic(err)
		}
	}(w)

	_, err = io.Copy(w, inFile)
	if err != nil {
		return "", err
	}

	if err := os.Remove(inputPath); err != nil {
		return "", fmt.Errorf("compression succeeded, but failed to delete original: %w", err)
	}

	return outputPath, err
}
