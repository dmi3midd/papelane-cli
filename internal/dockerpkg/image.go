package dockerpkg

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/viper"
)

func IsImageInstalled() bool {
	image := viper.GetString("image")
	cmd := exec.Command("docker", "image", "inspect", image)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func PullImage() {
	image := viper.GetString("image")
	cmd := exec.Command("docker", "pull", image)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("error while pulling image: %v\n", err)
		return
	}
	fmt.Println("Image pulled successfully.")
}
