/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package root

import (
	"fmt"

	"github.com/henrywhitaker3/srep/internal/cmd/check"
	"github.com/henrywhitaker3/srep/internal/cmd/list"
	"github.com/henrywhitaker3/srep/internal/cmd/run"
	"github.com/spf13/cobra"
)

func BuildRootCmd(version, commit, date string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "srep",
		Short:   "SRE practice",
		Long:    `A CLI that runs SRE practice scenarios`,
		Version: version,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cmd.SetVersionTemplate(fmt.Sprintf("%s version %s commit %s built at %s\n", cmd.Use, version, commit, date))

	cmd.AddCommand(list.NewListCommand())
	cmd.AddCommand(run.NewRunCommand())
	cmd.AddCommand(check.NewCheckCommand())

	return cmd
}
