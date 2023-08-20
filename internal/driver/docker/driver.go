package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/henrywhitaker3/srep/internal/driver"
	"github.com/henrywhitaker3/srep/internal/metadata"
)

type Docker struct {
	client client.APIClient
}

func NewDockerDriver() (*Docker, error) {
	d := &Docker{}

	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	d.client = c

	return d, nil
}

func (d *Docker) Create(s metadata.Scenario) (driver.Instance, error) {
	instance := Container{
		Name:  s.Name,
		Image: fmt.Sprintf("%s%s:%s", driver.ImagePrefix, s.Name, s.Version),
	}

	return &instance, nil
}

func (d *Docker) Run(ctx context.Context, i driver.Instance) error {
	c := i.(*Container)

	// d.pullImage(ctx, c.Image)

	return d.createContainer(ctx, c)
}

func (d *Docker) ConnectionCommand(i driver.Instance) string {
	c := i.(*Container)

	return fmt.Sprintf("docker exec -it %s bash", c.Name)
}

func (d *Docker) Kill(ctx context.Context, i driver.Instance) error {
	c := i.(*Container)
	err := d.client.ContainerStop(ctx, c.Name, container.StopOptions{})
	if err != nil {
		return err
	}
	return d.client.ContainerRemove(ctx, c.Name, types.ContainerRemoveOptions{})
}

func (d *Docker) Check(ctx context.Context, i driver.Instance) bool {
	c := i.(*Container)
	respID, err := d.client.ContainerExecCreate(ctx, c.Name, types.ExecConfig{
		Cmd:          []string{"/opt/check.sh"},
		AttachStdout: true,
	})
	if err != nil {
		return false
	}

	resp, err := d.client.ContainerExecAttach(ctx, respID.ID, types.ExecStartCheck{})
	if err != nil {
		return false
	}
	defer resp.Close()

	out := new(strings.Builder)
	if _, err := io.Copy(out, resp.Reader); err != nil {
		return false
	}

	// Check it is equal to start, nil *6, text start OK
	return out.String() == string([]byte{1, 0, 0, 0, 0, 0, 0, 2, 79, 75})
}

func (d *Docker) pullImage(ctx context.Context, image string) error {
	_, err := d.client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (d *Docker) createContainer(ctx context.Context, c *Container) error {
	resp, err := d.client.ContainerCreate(ctx, &container.Config{
		Image: c.Image,
	}, nil, nil, nil, c.Name)
	if err != nil {
		return err
	}
	c.Id = resp.ID

	return d.client.ContainerStart(ctx, c.Id, types.ContainerStartOptions{})
}
