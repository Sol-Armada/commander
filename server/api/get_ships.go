package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/ship"
)

func (s Server) GetApiV1Ships(c echo.Context) error {
	return c.JSON(200, ApiResponse{
		Data: ship.GetShips(),
	})
}
