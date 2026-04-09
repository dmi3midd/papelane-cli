package commands

import (
	"papelane-cli/internal/dockerpkg"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the docker is ready",
	Long:  `Check if the docker is ready by sending a request to the Telegram Bot API (Local) endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		dockerpkg.RunDockerPipeline()
	},
}
