package app

import (
	"github.com/jimfurnier/itemGenerator/internal/config"
	"github.com/jimfurnier/itemGenerator/internal/generator"
	"github.com/jimfurnier/itemGenerator/internal/writer"
)

func Execute(cfg *config.Config) (*writer.WriteResult, error) {
	// Create writer spec from config
	spec, err := writer.NewWriteSpec(cfg)
	if err != nil {
		return nil, err
	}

	// Build file writer
	fileWriter := writer.NewCompressedFileWriter(
		writer.NewDefaultFileWriter(generator.New(cfg)),
	)

	// Generate file
	result, err := fileWriter.Write(spec)
	if err != nil {
		return nil, err
	}

	return result, nil
}
