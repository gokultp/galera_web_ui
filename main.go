package main

import (
	"log"

	"github.com/gokultp/galera_web_ui/apis"
)

func main() {
	api, err := apis.NewAPI()
	if err != nil {
		log.Fatal(err)
	}
	api.Listen(":80")

}
