package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.HTML(http.StatusOK, "Hello, world")
}
