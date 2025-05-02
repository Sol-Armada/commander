package wsmanager

import (
	"log/slog"
	"time"

	"github.com/rs/xid"
	"github.com/sol-armada/commander/auth"
	ws "golang.org/x/net/websocket"
)

type Client struct {
	Id       string
	Send     chan Message
	LastSeen time.Time
	Active   bool
	MemberId string

	conn *ws.Conn
}

func (h *Hub) NewClient(conn *ws.Conn) *Client {
	client := &Client{
		Id:       xid.New().String(),
		Send:     make(chan Message),
		LastSeen: time.Now(),
		Active:   true,

		conn: conn,
	}
	go client.handleMessages(h)
	return client
}

func (c *Client) handleMessages(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
		c.Active = false
	}()

	for {
		var message Message
		if err := ws.JSON.Receive(c.conn, &message); err != nil {
			if err.Error() == "EOF" {
				slog.Debug("Connection closed by client", "clientId", c.Id)
				return
			}

			slog.Info("Error receiving message", "error", err)
			break
		}

		switch message.Channel {
		case ChannelRegister:
			h.logger.Info("Received message on register channel", "message", message)
			tokenRaw := message.Data["token"].(string)
			token, err := auth.GetToken(tokenRaw)
			if err != nil {
				slog.Error("Error getting token", "error", err)
				break
			}
			c.MemberId = token.Claims.(*auth.JwtClaims).Member.Id
			h.register <- c

			// sendMessage := Message{
			// 	Type:    MessageTypeRegistered,
			// 	Data:    map[string]any{"id": c.Id},
			// 	Channel: ChannelRegister,
			// }

			// if err := ws.JSON.Send(c.conn, sendMessage); err != nil {
			// 	slog.Error("Error sending message", "error", err)
			// }

		case ChannelKeepAlive:
			if err := keepAliveChannelHandler(h, c, message); err != nil {
				slog.Error("Error handling keep-alive channel message", "error", err)
			}
		case ChannelOperations:
			h.logger.Info("Received message on operations channel", "message", message)
			if err := operationsChannelHandler(h, c, message); err != nil {
				slog.Error("Error handling operations channel message", "error", err)
			}
		case ChannelMembers:
			h.logger.Info("Received message on members channel", "message", message)
			if err := membersChannelHandler(h, c, message); err != nil {
				slog.Error("Error handling members channel message", "error", err)
			}
		}
	}
}
