package main

import (
	"log"

	"github.com/dmi3midd/papelane-cli/internal/commands"
	"github.com/dmi3midd/papelane-cli/internal/logger"
)

func main() {
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	rootCmd := commands.RootCmd
	commands.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Log.Error("Error while execute root cmd", "error", err)
	}
}
