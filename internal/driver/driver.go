package driver

import "github.com/henrywhitaker3/srep/internal/metadata"

type Instance interface {
	// Get a shell into the running instance
	Shell() error
	// Kill the instance and clean it up
	Kill() error
}

type Driver interface {
	// Create a new instance of the scenario
	Create(metadata.Scenario) (*Instance, error)
}
