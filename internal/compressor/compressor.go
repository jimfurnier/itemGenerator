package compressor

import "fmt"

type Compressor interface {
	Compress(inputPath string) (string, error)
}

func GetCompressor(name string) (Compressor, error) {
	switch name {
	case "gzip":
		return &GzipCompressor{}, nil
	case "zip":
		return &ZipCompressor{}, nil
	default:
		return nil, fmt.Errorf("invalid compressor: %s", name)
	}
}
