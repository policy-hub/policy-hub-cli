package main

import (
	"os"

	"github.com/policy-hub/policy-hub-cli/internal/commands"
)

func main() {
	if err := commands.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
