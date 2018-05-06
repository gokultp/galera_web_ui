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
	err = api.Listen(":8000")
	if err != nil {
		fmt.Println(err)
	}

}
