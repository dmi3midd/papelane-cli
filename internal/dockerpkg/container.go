package dockerpkg

import (
	"fmt"
	"os/exec"

	"github.com/dmi3midd/papelane-cli/internal/config"
)

func DoesContainerExist() bool {
	containerName := config.GlobalConfig.GetString("containerName")
	cmd := exec.Command("docker", "inspect", containerName)
	return cmd.Run() == nil
}

func RunDockerContainer() error {
	port := config.GlobalConfig.GetInt("port")
	containerName := config.GlobalConfig.GetString("containerName")
	volume := config.GlobalConfig.GetString("volume")
	apiId := config.GlobalConfig.GetString("apiId")
	apiHash := config.GlobalConfig.GetString("apiHash")
	image := config.GlobalConfig.GetString("image")
	args := []string{
		"run", "-d",
		"-p", fmt.Sprintf("%d:8081", port),
		"--name", containerName,
		"--restart", "always",
		"-v", fmt.Sprintf("%s:/var/lib/telegram-bot-api", volume),
		"-e", fmt.Sprintf("TELEGRAM_API_ID=%s", apiId),
		"-e", fmt.Sprintf("TELEGRAM_API_HASH=%s", apiHash),
		image,
	}

	cmd := exec.Command("docker", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run container: %v", err)
	}
	return nil
}
