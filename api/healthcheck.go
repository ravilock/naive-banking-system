package api

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (s *server) healthcheck(c echo.Context) error {
	if err := s.db.Pool.Ping(); err != nil {
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintln("OK"))
}
