package main

import (
	"log"

	"github.com/dmi3midd/papelane-cli/internal/commands"
)

func main() {
	rootCmd := commands.RootCmd
	commands.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error while execute root cmd: %v", err)
	}
}
