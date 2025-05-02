package wsmanager

import (
	"github.com/labstack/echo/v4"
)

func GetWsShips(c echo.Context) (errOut error) {
	// ws.Handler(func(w *ws.Conn) {
	// 	defer w.Close()

	// 	token := ""
	// 	if err := ws.Message.Receive(w, &token); err != nil {
	// 		if err.Error() == "EOF" {
	// 			c.Logger().Debug("Connection closed by client")
	// 			return
	// 		}
	// 		errOut = errors.Join(err, errors.New("failed to receive message"))
	// 		return
	// 	}

	// 	if _, err := auth.GetToken(token); err != nil {
	// 		errOut = errors.Join(err, errors.New("failed to check token claims"))
	// 		return
	// 	}

	// 	authToken, err := auth.GetToken(token)
	// 	if err != nil {
	// 		errOut = errors.Join(err, errors.New("failed to check token claims"))
	// 		return
	// 	}

	// 	id := authToken.Claims.(*auth.JwtClaims).Member.Id
	// 	wsmanager.AddConnection(id, "ships", w)

	// 	for {
	// 		select {
	// 		case <-c.Request().Context().Done():
	// 			c.Logger().Debug("Connection closed by client")
	// 			return
	// 		default:
	// 		}

	// 		msg := ""
	// 		if err := ws.Message.Receive(w, &msg); err != nil {
	// 			if err.Error() == "EOF" {
	// 				c.Logger().Debug("Connection closed by client")
	// 				break
	// 			}
	// 			c.Logger().Error(err)
	// 		}
	// 		fmt.Printf("%s\n", msg)

	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }).ServeHTTP(c.Response(), c.Request())

	return nil
}
