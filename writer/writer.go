package writer

import (
	"github.com/jimfurnier/itemGenerator/config"
)

type Writer interface {
	WriteHeader([]string) error
	WriteRow([]string) error
	Close() error
}

func NewWriter(cfg *config.Config) (Writer, error) {
	switch cfg.Delimiter {
	case "csv":
		return NewCSVWriter(cfg.OutputName + ".csv")
	case "tab":
		return NewTSVWriter(cfg.OutputName +".tsv")
	default:
		return nil, nil
	}
}

