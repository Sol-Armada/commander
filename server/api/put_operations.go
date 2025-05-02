package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) PutApiV1OperationsOperationId(c echo.Context, id string) error {
	// get the body
	body := c.Request().Body
	defer body.Close()

	// read the body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to read request body")
	}

	op := &Operation{}
	if err := json.Unmarshal(bodyBytes, op); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to unmarshal request body")
	}

	return nil
}
