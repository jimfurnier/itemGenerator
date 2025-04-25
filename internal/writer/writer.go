package writer

import "fmt"

type Writer interface {
	WriteHeader([]string) error
	WriteRow([]string) error
	Close() error
}

func NewWriter(spec *WriteSpec) (Writer, error) {
	switch spec.name {
	case "comma":
		return NewDelimitedWriter(spec, supportedDelimiters["comma"].Char)
	case "tab":
		return NewDelimitedWriter(spec, supportedDelimiters["tab"].Char)
	default:
		return nil, fmt.Errorf("invalid writer type used: %s", spec.name)
	}
}
