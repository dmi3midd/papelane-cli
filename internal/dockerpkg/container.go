package dockerpkg

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func DoesContainerExist() bool {
	containerName := viper.GetString("containerName")
	cmd := exec.Command("docker", "inspect", containerName)
	return cmd.Run() == nil
}

func RunDockerContainer() error {
	port := viper.GetInt("port")
	containerName := viper.GetString("containerName")
	volume := viper.GetString("volume")
	apiId := viper.GetString("apiId")
	apiHash := viper.GetString("apiHash")
	image := viper.GetString("image")
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
		return fmt.Errorf("error while running container: %v", err)
	}
	fmt.Println("Container run successfully.")
	return nil
}
