package compressor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type ZipCompressor struct{}

func (ZipCompressor) Compress(inputPath, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	fileToZip, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(inputPath)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

