package config

type Config struct {
	outputName  string
	compression string
	delimiter   string
	columns     []Column
	rows        int
}

func (c *Config) GetHeaders() []string {
	headers := make([]string, len(c.columns))
	for i, col := range c.columns {
		headers[i] = col.Name
	}
	return headers
}

func (c *Config) GetColumns() []Column {
	return c.columns
}

func (c *Config) GetDelimiter() string {
	return c.delimiter
}

func (c *Config) GetCompression() string {
	return c.compression
}

func (c *Config) GetOutputName() string {
	return c.outputName
}

func (c *Config) GetRows() int {
	return c.rows
}

func (c *Config) ForceCompression(compression string) {
	c.compression = compression
}
