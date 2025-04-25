package writer

import (
	"encoding/csv"
	"os"
)

type writer struct {
	file    *os.File
	writer  *csv.Writer
}

func NewCSVWriter(path string) (*writer, error) {
	return NewWriterWithDelimiter(path, ',')
}

func NewTSVWriter(path string) (*writer, error) {
	return NewWriterWithDelimiter(path, '\t')
}


func NewWriterWithDelimiter(path string, delimiter rune) (*writer, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	w := csv.NewWriter(f)
	w.Comma = delimiter

	return &writer{f, w}, nil
}

func (w *writer) WriteHeader(headers []string) error {
	return w.writer.Write(headers)
}

func (w *writer) WriteRow(row []string) error {
	return w.writer.Write(row)
}

func (w *writer) Close() error {
	w.writer.Flush()
	return w.file.Close()
}

