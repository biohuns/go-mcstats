package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/andrewtian/minepong"
)

func main() {
	var (
		host = flag.String("host", "", "hostname string")
		port = flag.Int("port", 25565, "port number")
	)
	flag.Parse()

	if *host == "" {
		fmt.Fprintln(os.Stderr, fmt.Errorf("host is required"))
		os.Exit(1)
	}
	stats, err := getStats(fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	jsonString, err := json.Marshal(stats)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(jsonString))
}

func getStats(host string) (*minepong.Pong, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	stats, err := minepong.Ping(conn, host)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
