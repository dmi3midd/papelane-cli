package commands

import (
	"fmt"

	"papelane-cli/internal/config"
	"papelane-cli/internal/database"
	"papelane-cli/internal/domain"
	"papelane-cli/internal/repositories"
	"papelane-cli/internal/telegrampkg"

	"github.com/spf13/cobra"
)

var (
	client     telegrampkg.Client
	db         database.Service
	folderRepo domain.FolderRepository
	fileRepo   domain.FileRepository
)

var RootCmd = &cobra.Command{
	Use:   "papelane",
	Short: "Papelane-CLI is a tool that turns your Telegram bot into cloud storage.",
	Long: `Papelane-CLI is a high-performance CLI utility
		that transforms your Telegram bot into a structured cloud storage system
		by leveraging a local SQLite database for instant file navigation 
		and the Telegram Bot API (Local) to bypass standard 50MB upload limits. `,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "init" {
			return nil
		}
		if err := config.ReadInGlobalCfg(); err != nil {
			fmt.Printf("Error reading global config: %v\n", err)
			return err
		}
		if err := config.ReadInCurrDirCfg(); err != nil {
			fmt.Printf("Error reading current directory config: %v\n", err)
			return err
		}

		// Initialize database and repositories
		dbPath := config.GlobalConfig.GetString("dbPath")
		if dbPath != "" {
			db = database.New(dbPath)
		}
		folderRepo = repositories.NewFolderRepository(db.GetDB())
		fileRepo = repositories.NewFileRepository(db.GetDB())

		// Initialize Telegram client
		botToken := config.GlobalConfig.GetString("botToken")
		port := config.GlobalConfig.GetInt("port")
		var err error
		if botToken != "" && port != 0 {
			client, err = telegrampkg.NewTelegramClient(botToken, fmt.Sprintf("http://localhost:%d", port))
			if err != nil {
				fmt.Printf("Error creating Telegram client: %v\n", err)
				return err
			}
		}

		return nil
	},
}

var goToCurrDirFlag bool
var filesFlag bool
var dirsFlag bool

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(currCmd)
	rootCmd.AddCommand(toRootCmd)
	rootCmd.AddCommand(mkdirCmd)
	rootCmd.AddCommand(cdCmd)
	rootCmd.AddCommand(rmdCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(uploadCmd)

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

	mkdirCmd.Flags().BoolVarP(&goToCurrDirFlag, "cd", "d", false, "Go to the newly created directory")

	lsCmd.Flags().BoolVarP(&filesFlag, "files", "f", false, "List files in the current directory")
	lsCmd.Flags().BoolVarP(&dirsFlag, "dirs", "d", false, "List directories in the current directory")
}
