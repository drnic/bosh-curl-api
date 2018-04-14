package main

import (
	"fmt"

	"github.com/drnic/bosh-curl-example/boshcli"
)

func main() {
	boshcli.Check()

	fmt.Printf("%v\n", boshcli.GetDeployments())
}
