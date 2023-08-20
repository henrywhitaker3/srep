package docker

import (
	"github.com/henrywhitaker3/srep/internal/driver"
	"github.com/henrywhitaker3/srep/internal/metadata"
)

type Docker struct{}

func (d *Docker) Create(metadata.Scenario) (*driver.Instance, error) {
	return nil, nil
}
