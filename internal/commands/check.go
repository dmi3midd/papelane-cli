package commands

import (
	"fmt"

	"github.com/dmi3midd/papelane-cli/internal/dockerpkg"

	"github.com/spf13/cobra"
)

// check command for checking if the docker is ready
// example: check
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the docker is ready",
	Long:  `Check if the docker is ready by sending a request to the Telegram Bot API (Local) endpoint.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := dockerpkg.RunDockerPipeline(); err != nil {
			return fmt.Errorf("failed to run docker pipeline: %w", err)
		}
		return nil
	},
}
