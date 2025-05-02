package api

import (
	"time"

	"github.com/labstack/echo/v4"
	ws "golang.org/x/net/websocket"
)

func (s Server) GetWs(ctx echo.Context) error {
	ws.Handler(func(ws *ws.Conn) {
		c := s.wm.NewClient(ws)

		for c.Active {
			time.Sleep(1 * time.Second)
		}
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}
