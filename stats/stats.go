package stats

import (
	"net"

	"github.com/andrewtian/minepong"
)

func GetStats(host string) (*minepong.Pong, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	pong, err := minepong.Ping(conn, host)
	if err != nil {
		return nil, err
	}
	return pong, nil
}
