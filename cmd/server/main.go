package main

import (
	"fmt"
	"log"

	"github.com/thiagohmm/ApiGo/config"
)

func main() {
	//load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Database running on port", config.DBPORT)

}
