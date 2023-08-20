package docker

type Container struct {
	Id string
}

func (c *Container) Shell() error {
	return nil
}

func (c *Container) Kill() error {
	return nil
}
