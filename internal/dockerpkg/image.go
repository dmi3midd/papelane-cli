package dockerpkg

import (
	"log"
	"os/exec"

	"github.com/dmi3midd/papelane-cli/internal/config"
)

func IsImageInstalled() bool {
	image := config.GlobalConfig.GetString("image")
	cmd := exec.Command("docker", "image", "inspect", image)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func PullImage() {
	image := config.GlobalConfig.GetString("image")
	cmd := exec.Command("docker", "pull", image)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("failed to pull image: %v\n", err)
		return
	}
}
