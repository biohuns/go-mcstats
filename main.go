package main

import (
	"log"

	"github.com/biohuns/go-minecraft-stats/config"
	"github.com/biohuns/go-minecraft-stats/server"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	s := server.NewServer(cfg)
	s.Start()
}
