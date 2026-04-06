package main

import (
	"log"
	"papelane-cli/internal/commands"
)

func main() {
	log.Println("papelane-cli is running")
	rootCmd := commands.RootCmd
	commands.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error while execute root cmd: %v", err)
	}
}
