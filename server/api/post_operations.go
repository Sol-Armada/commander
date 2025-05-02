package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/member"
	"github.com/sol-armada/commander/operation"
)

func (s Server) PostApiV1Operations(c echo.Context) error {
	id := c.Get("member").(member.Member).Id

	op := operation.NewOperation(id, "")
	if err := s.wm.OperationsCreate(op); err != nil {
		return c.JSON(500, ApiResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(200, ApiResponse{
		Data: op,
	})
}
