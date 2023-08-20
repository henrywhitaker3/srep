package docker

import "github.com/henrywhitaker3/srep/internal/metadata"

type Container struct {
	Id    string
	Name  string
	Image string
	Ports []metadata.Port
}
