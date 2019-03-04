package server

import (
	"fmt"
	"net/http"

	"github.com/biohuns/go-minecraft-stats/stats"
	"github.com/labstack/echo"
)

func statsHandler(c echo.Context) error {
	cc := c.(*configuredContext)
	host := fmt.Sprintf(
		"%s:%d",
		cc.config.MC.Host,
		cc.config.MC.Port,
	)
	// get stats
	stats, err := stats.GetStats(host)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, stats)
}
