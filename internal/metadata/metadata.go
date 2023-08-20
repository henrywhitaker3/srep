package metadata

import (
	_ "embed"
	"encoding/json"
)

//go:embed metadata.json
var mdJson string

type Metadata []Scenario

type Scenario struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Version     string `json:"version"`
}

func Get() (*Metadata, error) {
	md := &Metadata{}

	if err := json.Unmarshal([]byte(mdJson), md); err != nil {
		return nil, err
	}

	return md, nil
}
