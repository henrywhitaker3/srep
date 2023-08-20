/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package cmd

import (
	"github.com/henrywhitaker3/srep/internal/metadata"
	"github.com/henrywhitaker3/srep/internal/views/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available practice scenarios",
	RunE: func(cmd *cobra.Command, args []string) error {
		md, err := metadata.Get()
		if err != nil {
			return err
		}

		tbl := list.NewTable(md)
		tbl.Print()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
