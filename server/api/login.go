package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sol-armada/commander/auth"
)

func (s Server) GetApiV1Login(c echo.Context, params GetApiV1LoginParams) error {
	code := params.Code
	if code == "" {
		return errors.New("code is required")
	}

	token, err := auth.Login(s.DB(), code)
	if err != nil {
		if err.Error() == "invalid_grant" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"status":  "fail",
				"message": "invalid_grant",
			})
		}

		return err
	}

	return c.JSON(http.StatusOK, &ApiResponse{
		Data: map[string]any{
			"token": token,
		},
	})
}
