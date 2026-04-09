package telegrampkg

import (
	"fmt"
	"net/http"
)

type Client struct {
	botToken string
	baseURL  string
	client   *http.Client
}

func NewClient(botToken string, baseURL string) *Client {
	return &Client{
		botToken: botToken,
		baseURL:  baseURL,
		client:   &http.Client{},
	}
}

func (c *Client) Ping() error {
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
