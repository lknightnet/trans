package main

import (
	"log"
	"transcription/config"
	"transcription/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println()
	}

	app.Run(cfg)
}
