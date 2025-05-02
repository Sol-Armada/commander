package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/member"
)

func (s Server) GetApiV1Members(c echo.Context) error {
	members := member.List(s.DB())
	return c.JSON(http.StatusOK, &ApiResponse{
		Data: members,
	})
}
