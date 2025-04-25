package writer

import (
	"encoding/csv"
	"os"
)

type DelimitedWriter struct {
	file   *os.File
	writer *csv.Writer
}

func NewDelimitedWriter(spec *WriteSpec, delimiter rune) (*DelimitedWriter, error) {
	f, err := os.Create(spec.path)
	if err != nil {
		return nil, err
	}

	w := csv.NewWriter(f)
	w.Comma = delimiter

	return &DelimitedWriter{f, w}, nil
}

func (w *DelimitedWriter) WriteHeader(headers []string) error {
	return w.writer.Write(headers)
}

func (w *DelimitedWriter) WriteRow(row []string) error {
	return w.writer.Write(row)
}

func (w *DelimitedWriter) Close() error {
	w.writer.Flush()
	return w.file.Close()
}
