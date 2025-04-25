package writer

import (
	"github.com/jimfurnier/itemGenerator/internal/generator"
)

type FileWriter interface {
	Write(spec *WriteSpec) (*WriteResult, error)
}

type DefaultFileWriter struct {
	generator *generator.Generator
}

func NewDefaultFileWriter(generator *generator.Generator) *DefaultFileWriter {
	return &DefaultFileWriter{generator}
}

func (fileWriter *DefaultFileWriter) Write(spec *WriteSpec) (*WriteResult, error) {
	// Create writer and defer closing
	writer, err := NewWriter(spec)
	if err != nil {
		return nil, err
	}
	defer func(writer Writer) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	// Write the headers
	err = writer.WriteHeader(spec.headers)
	if err != nil {
		return nil, err
	}

	// Write the rows
	for i := 0; i < spec.count; i++ {
		row := fileWriter.generator.GenerateRow(i)
		err := writer.WriteRow(row)
		if err != nil {
			return nil, err
		}
	}

	return NewResult(spec.path, spec.count), nil
}
