/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package run

import (
	"fmt"

	"github.com/henrywhitaker3/srep/internal/driver"
	"github.com/henrywhitaker3/srep/internal/driver/docker"
	"github.com/henrywhitaker3/srep/internal/metadata"
	"github.com/spf13/cobra"
)

var (
	k8s bool
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [scenario]",
		Short: "Run the specified practice scenarios",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := metadata.Find(args[0])
			if err != nil {
				return err
			}
			fmt.Println(s)

			d := getDriver()
			d.Create(*s)

			return nil
		},
	}

	cmd.Flags().BoolVar(&k8s, "k8s", false, "Determine whether to use kubernetes as the driver. Defaults to using docker.")

	return cmd
}

func getDriver() driver.Driver {
	if k8s {

	}
	return &docker.Docker{}
}
