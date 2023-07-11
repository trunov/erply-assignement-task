package main

import (
	"log"

	"github.com/trunov/erply-assignement-task/user-service/internal/app"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
