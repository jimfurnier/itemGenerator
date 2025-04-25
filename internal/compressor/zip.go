package compressor

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ZipCompressor struct{}

func (ZipCompressor) Compress(inputPath string) (string, error) {
	outputPath := inputPath + ".zip"

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

	zipWriter := zip.NewWriter(outFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(zipWriter)

	fileToZip, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			panic(err)
		}
	}(fileToZip)

	info, err := fileToZip.Stat()
	if err != nil {
		return "", err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return "", err
	}
	header.Name = filepath.Base(inputPath)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return "", err
	}

	if err := os.Remove(inputPath); err != nil {
		return "", fmt.Errorf("compression succeeded, but failed to delete original: %w", err)
	}

	return outputPath, err
}
