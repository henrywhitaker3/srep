/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package describe

import (
	"fmt"

	"github.com/henrywhitaker3/srep/internal/metadata"
	"github.com/spf13/cobra"
)

func NewDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe [scenario]",
		Short: "Describe the specified practice scenario",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := metadata.Find(args[0])
			if err != nil {
				return err
			}

			fmt.Printf("Scenario: %s\n", s.Name)
			fmt.Printf("Difficulty: %s\n", s.Difficulty)
			fmt.Println("Tags:")
			for _, tag := range s.Tags {
				fmt.Printf("  - %s\n", tag)
			}
			fmt.Printf("\nDescription:\n")
			fmt.Println(s.Description)

			return nil
		},
	}

	return cmd
}
