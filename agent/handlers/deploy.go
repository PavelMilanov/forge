package handlers

import (
	"github.com/labstack/echo/v4"
)

func deployHandler(c echo.Context) error {
	return c.String(200, "s")
}
