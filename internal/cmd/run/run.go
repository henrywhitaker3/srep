/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package run

import (
	"fmt"

	"github.com/henrywhitaker3/srep/internal/cmd/common"
	"github.com/henrywhitaker3/srep/internal/metadata"
	"github.com/spf13/cobra"
)

var (
	k8s bool
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "run [scenario]",
		Short:     "Run the specified practice scenarios",
		Args:      cobra.ExactArgs(1),
		ValidArgs: common.ScenarioCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := metadata.Find(args[0])
			if err != nil {
				return err
			}

			d, err := common.GetDriver(k8s)
			if err != nil {
				return err
			}
			instance, err := d.Create(*s)
			if err != nil {
				return err
			}
			if err := d.Run(cmd.Context(), instance); err != nil {
				return err
			}

			fmt.Printf("To connect to the instance, run the following command:\n\n")
			fmt.Println(d.ConnectionCommand(instance))
			fmt.Printf("\nWhen you have finished, you can run `srep check %s` to check your work\n", args[0])

			return nil
		},
	}

	cmd.Flags().BoolVar(&k8s, "k8s", false, "Determine whether to use kubernetes as the driver. Defaults to using docker.")

	return cmd
}
