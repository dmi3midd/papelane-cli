package dockerpkg

import "os/exec"

func IsDockerInstalled() bool {
	cmd := exec.Command("docker", "-v")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
