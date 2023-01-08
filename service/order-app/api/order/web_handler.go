package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ss *service) fetch(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
