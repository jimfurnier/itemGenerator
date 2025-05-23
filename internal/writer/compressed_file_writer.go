package writer

import (
	"fmt"
	"github.com/jimfurnier/itemGenerator/internal/compressor"
)

type CompressedFileWriter struct {
	fw FileWriter
}

func NewCompressedFileWriter(fw FileWriter) *CompressedFileWriter {
	return &CompressedFileWriter{fw}
}

func (f *CompressedFileWriter) Write(spec *WriteSpec) (*WriteResult, error) {
	result, err := f.fw.Write(spec)
	if err != nil {
		return nil, err
	}

	switch spec.Compression() {
	case "zip", "gzip":
		comp, err := compressor.GetCompressor(spec.Compression())
		if err != nil {
			return nil, err
		}

		newPath, err := comp.Compress(result.Path())
		if err != nil {
			return nil, err
		}

		return NewResult(newPath, spec.count), nil
	case "none":
		return result, nil
	default:
		return nil, fmt.Errorf("invalid compression type used: %s", spec.Compression())
	}
}
