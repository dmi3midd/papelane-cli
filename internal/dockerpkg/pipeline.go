package dockerpkg

import "fmt"

func RunDockerPipeline() error {
	cond := IsDockerInstalled()
	if !cond {
		return fmt.Errorf("docker is not installed")
	}
	cond = IsImageInstalled()
	if !cond {
		PullImage()
	}
	cond = DoesContainerExist()
	if !cond {
		RunDockerContainer()
	}
	return nil
}
