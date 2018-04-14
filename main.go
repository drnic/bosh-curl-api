package main

import (
	"fmt"

	"github.com/drnic/bosh-curl-api/boshcli"
)

func main() {
	boshcli.Check()

	deployments := boshcli.GetDeployments()
	fmt.Printf("%v\n", deployments)

	for _, deployment := range *deployments {
		deploymentManifest := boshcli.GetDeploymentManifest(deployment.Name)
		fmt.Println(deploymentManifest.Manifest)
	}
}
