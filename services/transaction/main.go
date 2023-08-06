package main

import (
	"github.com/hoyirul/go-microservices-api/config"
)

func main() {
	config.ConnectDB("TRANSACTION")
}
