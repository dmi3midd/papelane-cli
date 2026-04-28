package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

var Log *slog.Logger

func InitLogger() error {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config dir: %w", err)
	}

	appDir := filepath.Join(cfgDir, "papelane-cli")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("failed to create app config dir: %w", err)
	}

	logFilePath := filepath.Join(appDir, "papelane.log")

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	Log = slog.New(handler)
	slog.SetDefault(Log)

	return nil
}
