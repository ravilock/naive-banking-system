package api

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintln("OK"))
}
