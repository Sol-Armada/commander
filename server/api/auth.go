package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) GetApiV1Auth(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
