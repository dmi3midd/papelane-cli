package telegrampkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Ping() {
	url := fmt.Sprintf("http://localhost:%v/bot%s/getMe", viper.GetInt("port"), viper.GetString("botToken"))

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to connect to Telegram Bot API: %v", err)
		return
	}
	log.Printf("Telegram Bot API is ready (status: %d)", resp.StatusCode)
	resp.Body.Close()
}
