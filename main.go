package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// check that 'bosh curl' available
	cmd := exec.Command("sh", "-c", "bosh curl -h")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal("Need 'bosh curl' from https://github.com/cloudfoundry/bosh-cli/pull/408")
	}

	// check that bosh environment configured and connectable
	cmd = exec.Command("sh", "-c", "bosh curl /info")
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal("Cannot connect to BOSH")
	}

	fmt.Printf("%s\n", stdoutStderr)
}
