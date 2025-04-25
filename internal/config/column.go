package config

import "encoding/json"

const OptionalDefault = 100
const TypeDefault = "string"

type Column struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Optional int      `json:"optional"`
	Random   []string `json:"random"`
}

func (c *Column) UnmarshalJSON(data []byte) error {
	type Alias Column

	aux := &struct {
		Type     *string `json:"type"`
		Optional *int    `json:"optional"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Apply defaults
	if aux.Type == nil {
		c.Type = TypeDefault
	} else {
		c.Type = *aux.Type
	}

	if aux.Optional == nil {
		c.Optional = OptionalDefault
	} else {
		c.Optional = *aux.Optional
	}

	return nil
}
