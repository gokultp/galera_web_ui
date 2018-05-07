package main

import (
	"fmt"

	"github.com/gokultp/galera_web_ui/apis"
)

func main() {
	api, err := apis.NewAPI()
	if err != nil {
		fmt.Println(err)
	}
	err = api.Listen(":80")
	if err != nil {
		fmt.Println(err)
	}

}
