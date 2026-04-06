package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "papelane",
	Short: "Papelane-CLI is a tool that turns your Telegram bot into cloud storage.",
	Long: `Papelane-CLI is a high-performance CLI utility
		that transforms your Telegram bot into a structured cloud storage system
		by leveraging a local SQLite database for instant file navigation 
		and the Telegram Bot API (Local) to bypass standard 50MB upload limits. 
		It provides a persistent REPL environment where you can manage files using familiar shell commands like 'cd', 'ls', and 'mkdir', 
		supports multiple storage profiles for different bots, 
		and ensures efficient, memory-friendly data streaming of files up to 2GB directly through a local Docker-managed proxy.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("apid", "", "Your TELEGRAM_API_ID")
	initCmd.Flags().String("apih", "", "Your TELEGRAM_API_HASH")
	initCmd.Flags().String("token", "", "Your bot token")
	initCmd.Flags().Int("cid", 0, "Your chat id")
	initCmd.Flags().Int("port", 8081, "Post for docker container")
	initCmd.Flags().Bool("sa", false, "Always stop docker true or flase")

	initCmd.MarkFlagRequired("apid")
	initCmd.MarkFlagRequired("apih")
	initCmd.MarkFlagRequired("token")
	initCmd.MarkFlagRequired("cid")
}
