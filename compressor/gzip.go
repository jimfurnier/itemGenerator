package compressor

import (
	"compress/gzip"
	"io"
	"os"
)

type GzipCompressor struct{}

func (GzipCompressor) Compress(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := gzip.NewWriter(outFile)
	defer w.Close()
	_, err = io.Copy(w, inFile)
	return err
}

