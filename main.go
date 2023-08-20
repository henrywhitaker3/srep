/*
Copyright © 2023 Henry Whitaker <henrywhitaker3@outlook.com>
*/
package main

import (
	"os"

	"github.com/henrywhitaker3/srep/cmd/root"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	err := root.BuildRootCmd(version, commit, date).Execute()
	if err != nil {
		os.Exit(1)
	}
}
