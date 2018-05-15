package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gokultp/galera_web_ui/apis"
)

func main() {

	// getting docker daemon api version
	cmd := exec.Command("docker", "version", "--format", "{{.Server.APIVersion}}")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		fmt.Println("Something went wrong, start docker if you are not running.")
		panic(err)
	}
	apiVersion := strings.TrimSpace(string(cmdOutput.Bytes()))
	fmt.Println("supported api version", apiVersion)

	// setting env variable
	os.Setenv("DOCKER_API_VERSION", apiVersion)

	api, err := apis.NewAPI()
	if err != nil {
		panic(err)
	}
	err = api.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
