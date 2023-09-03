package common

import (
	"errors"

	"github.com/henrywhitaker3/srep/internal/driver"
	"github.com/henrywhitaker3/srep/internal/driver/docker"
	"github.com/henrywhitaker3/srep/internal/metadata"
)

func ScenarioCompletion() []string {
	out := []string{}
	s, err := metadata.Get()

	if err != nil {
		return out
	}

	for _, scenario := range *s {
		out = append(out, scenario.Name)
	}

	return out
}

func GetDriver(k8s bool) (driver.Driver, error) {
	if k8s {
		return nil, errors.New("k8s driver not implemented yet")
	}
	return docker.NewDockerDriver()
}
