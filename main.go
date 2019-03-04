package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/biohuns/go-minecraft-stats/config"
	"github.com/biohuns/go-minecraft-stats/stats"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	stats, err := stats.GetStats(fmt.Sprintf("%s:%s", cfg.Host, strconv.Itoa(cfg.Port)))
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("%+v\n", stats)
	jsonBytes, err := json.Marshal(stats)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
