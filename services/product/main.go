package main

import (
	"fmt"

	"github.com/hoyirul/go-microservices-api/config"
)

func main() {
	// config.ConnectDB("PRODUCT")
	config.LoadEnv()

	if err := config.LoadEnv(); err != nil {
		fmt.Println(err)
	}
}
