package config

import (
	"encoding/json"
	"os"
)

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Config struct {
	OutputName  string   `json:"output_name"`
	Compression string   `json:"compression"` // none, zip, gzip
	Delimiter   string   `json:"delimiter"`   // csv, tab
	Columns     []Column `json:"columns"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}

func (c *Config) GetHeaders() []string {
	headers := make([]string, len(c.Columns))
	for i, col := range c.Columns {
		headers[i] = col.Name
	}
	return headers
}

