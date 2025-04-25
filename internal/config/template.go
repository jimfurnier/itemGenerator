package config

import (
	"encoding/json"
	"os"
)

type rawConfig struct {
	OutputName  string   `json:"output_name"`
	Compression string   `json:"compression"` // none, zip, gzip
	Delimiter   string   `json:"delimiter"`   // comma, tab
	Columns     []Column `json:"columns"`
}

func LoadFromJsonTemplate(path string, rows int) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg rawConfig
	err = json.Unmarshal(data, &cfg)

	return &Config{
		cfg.OutputName,
		cfg.Compression,
		cfg.Delimiter,
		cfg.Columns,
		rows}, err
}
