package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/operation"
)

func (s Server) GetApiV1OperationsOperationId(c echo.Context, id string) error {
	op, err := operation.GetOperation(id)
	if err != nil {
		return c.JSON(404, ApiResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(200, ApiResponse{
		Data: op,
	})
}
