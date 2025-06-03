package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "basic.html", map[string]interface{}{})
}

func SecPage(c echo.Context) error {
	return c.Render(http.StatusOK, "basic.html", map[string]interface{}{})
}
