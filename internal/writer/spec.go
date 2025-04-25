package writer

import (
	"fmt"
	"github.com/jimfurnier/itemGenerator/internal/config"
)

const OutDirectory = "generated/"

type WriteSpec struct {
	name        string
	path        string
	compression string
	headers     []string
	count       int
}

type DelimiterType struct {
	Name      string
	Char      rune
	Extension string
}

var supportedDelimiters = map[string]DelimiterType{
	"comma": {Name: "comma", Char: ',', Extension: ".csv"},
	"tab":   {Name: "tab", Char: '\t', Extension: ".tsv"},
}

func NewWriteSpec(cfg *config.Config, rows *int) (*WriteSpec, error) {
	d, ok := supportedDelimiters[cfg.Delimiter]
	if !ok {
		return nil, fmt.Errorf("unsupported delimiter: %s", cfg.Delimiter)
	}

	return &WriteSpec{name: d.Name, path: OutDirectory + cfg.OutputName + d.Extension, headers: cfg.GetHeaders(), count: *rows, compression: cfg.Compression}, nil
}

func (spec *WriteSpec) Path() string {
	return spec.path
}

func (spec *WriteSpec) Compression() string {
	return spec.compression
}
