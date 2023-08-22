/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package check

import (
	"fmt"

	"github.com/henrywhitaker3/srep/internal/cmd"
	"github.com/henrywhitaker3/srep/internal/driver"
	"github.com/henrywhitaker3/srep/internal/driver/docker"
	"github.com/henrywhitaker3/srep/internal/metadata"
	"github.com/spf13/cobra"
)

var (
	k8s   bool
	clean bool
)

func NewCheckCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "check [scenario]",
		Short:     "Check the specified practice scenario",
		Args:      cobra.ExactArgs(1),
		ValidArgs: cmd.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := metadata.Find(args[0])
			if err != nil {
				return err
			}

			d, err := getDriver()
			if err != nil {
				return err
			}
			instance, err := d.Create(*s)
			if err != nil {
				return err
			}

			if d.Check(cmd.Context(), instance) {
				fmt.Println("The check script passed!")
				if clean {
					return d.Kill(cmd.Context(), instance)
				}
				return nil
			}

			fmt.Println("The check script failed, try again")

			return nil
		},
	}

	cmd.Flags().BoolVar(&k8s, "k8s", false, "Determine whether to use kubernetes as the driver. Defaults to using docker.")
	cmd.Flags().BoolVar(&clean, "clean", true, "Determine whether to kill and remove the instance.")

	return cmd
}

func getDriver() (driver.Driver, error) {
	if k8s {

	}
	return docker.NewDockerDriver()
}
