package telegrampkg

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"papelane-cli/internal/config"
	"strings"
	"time"

	"gopkg.in/telebot.v4"
)

type Client interface {
	Ping() error
	UploadFile(filePath string, fileInfo os.FileInfo) (*UploadedFile, error)
	DownloadFile(tgFileId string, destPath string) error
}

type TelegramClient struct {
	botToken string
	baseURL  string
	client   *http.Client
	bot      *telebot.Bot
}

func NewTelegramClient(botToken string, baseURL string) (*TelegramClient, error) {
	pref := telebot.Settings{
		URL:    baseURL,
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		Client: &http.Client{},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to create Telegram Bot: %v", err)
	}
	return &TelegramClient{
		botToken: botToken,
		baseURL:  baseURL,
		client:   &http.Client{},
		bot:      bot,
	}, nil
}

func (c *TelegramClient) Ping() error {
	url := fmt.Sprintf("%s/bot%s/getMe", c.baseURL, c.botToken)
	resp, err := c.client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to connect to Telegram Bot API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram Bot API returned status code %d", resp.StatusCode)
	}

	return nil
}

type UploadedFile struct {
	TgMsgId  int
	TgFileId string
	Name     string
	Size     int
	MimeType string
}

func (c *TelegramClient) UploadFile(filePath string, fileInfo os.FileInfo) (*UploadedFile, error) {
	// Open the file to read its content and detect MIME type
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buffer := make([]byte, 512)
	_, _ = f.Read(buffer)
	mimeType := http.DetectContentType(buffer)

	// Send the file to Telegram
	chatId := config.GlobalConfig.GetInt("chatId")
	recipient := &telebot.Chat{
		ID: int64(chatId),
	}
	pathParts := strings.Split(filePath, string(os.PathSeparator))
	document := &telebot.Document{
		FileName: pathParts[len(pathParts)-1],
		MIME:     mimeType,
		File:     telebot.FromDisk(filePath),
	}
	msg, err := c.bot.Send(recipient, document)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to Telegram: %v", err)
	}
	return &UploadedFile{
		TgMsgId:  msg.ID,
		TgFileId: document.FileID,
		Name:     pathParts[len(pathParts)-1],
		Size:     int(document.FileSize),
		MimeType: document.MIME,
	}, nil
}

func (c *TelegramClient) DownloadFile(tgFileId string, destPath string) error {
	file, err := c.bot.FileByID(tgFileId)
	if err != nil {
		return fmt.Errorf("failed to get file info (or download file to cache): %v", err)
	}

	containerName := config.GlobalConfig.GetString("containerName")

	// Get path to file inside container
	// Local Bot API without --local flag stores files in subfolders:
	// /var/lib/telegram-bot-api/<botToken>/<file.FilePath>
	containerFilePath := fmt.Sprintf("/var/lib/telegram-bot-api/%s/%s", c.botToken, file.FilePath)

	// Use docker cp for instant file copying,
	// this solves any problems with user access rights to the volume
	cmd := exec.Command("docker", "cp", fmt.Sprintf("%s:%s", containerName, containerFilePath), destPath)
	if err := cmd.Run(); err != nil {
		// HTTP download as a fallback (if the container is not available locally)
		err = c.bot.Download(&file, destPath)
		if err != nil {
			return fmt.Errorf("failed to copy (docker cp) and download file from Telegram: %v", err)
		}
	}
	return nil
}
