package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/operation"
)

func (s Server) GetApiV1Operations(c echo.Context) error {
	return c.JSON(200, ApiResponse{
		Data: operation.GetOperations(),
	})
}
