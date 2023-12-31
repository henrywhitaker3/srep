package driver

import (
	"context"

	"github.com/henrywhitaker3/srep/internal/metadata"
)

var (
	ImagePrefix = "ghcr.io/henrywhitaker3/srep/"
)

type Instance interface{}

type Driver interface {
	// Create a new instance of the scenario
	Create(metadata.Scenario) (Instance, error)
	// Run a new instance
	Run(context.Context, Instance) error
	// Get the connection command
	ConnectionCommand(Instance) string
	// Kill and remove the instance
	Kill(context.Context, Instance) error
	// Check the work in the instance
	Check(context.Context, Instance) bool
}
