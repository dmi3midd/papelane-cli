package dockerpkg

import "log"

func RunDockerPipeline() {
	cond := IsDockerInstalled()
	if !cond {
		log.Printf("Install docker first!")
		return
	}
	cond = IsImageInstalled()
	if !cond {
		PullImage()
	}
	cond = DoesContainerExist()
	if !cond {
		RunDockerContainer()
	}
}
