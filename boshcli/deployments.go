package boshcli

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Deployments []struct {
	Name     string `json:"name"`
	Releases []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"releases"`
	Stemcells []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"stemcells"`
	CloudConfig string        `json:"cloud_config"`
	Teams       []interface{} `json:"teams"`
}

// GetDeployments from target BOSH environment
func GetDeployments() (deployments *Deployments) {
	deployments = &Deployments{}
	cmd := exec.Command("sh", "-c", "bosh curl /deployments")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(stdoutStderr), deployments); err != nil {
		log.Fatal(err)
	}

	return
}
