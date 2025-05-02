package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/member"
)

func (s Server) GetApiV1MembersMemberId(c echo.Context, id string) error {
	var m *member.Member
	switch id {
	case "@me":
		me, ok := c.Get("member").(member.Member)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"status":  "error",
				"message": "unauthorized",
			})
		}
		m = &me
	default:
		var err error
		m, err = member.Get(s.DB(), id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"status":  "error",
				"message": "member not found",
			})
		}
	}

	if m == nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "error",
			"message": "member not found",
		})
	}

	return c.JSON(http.StatusOK, &ApiResponse{
		Data: m,
	})
}
