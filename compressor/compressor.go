package compressor

type Compressor interface {
	Compress(inputPath string, outputPath string) error
}

func GetCompressor(name string) (Compressor, error) {
	switch name {
	case "gzip":
		return &GzipCompressor{}, nil
	case "zip":
		return &ZipCompressor{}, nil
	default:
		return nil, nil
	}
}

