package wsmanager

import (
	"log/slog"
	"time"

	"github.com/sol-armada/sol-bot/stores"
	ws "golang.org/x/net/websocket"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client

	db *stores.Client

	logger *slog.Logger
}

func NewHub(logger *slog.Logger, db *stores.Client) *Hub {
	if logger == nil {
		logger = slog.Default()
	}

	logger = logger.With("component", "wsmanager")

	h := &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),

		db: db,

		logger: logger,
	}

	go h.KeepAlive()
	go func() {
		if err := h.watchForMemberUpdates(); err != nil {
			logger.Error("Error watching for member updates", "error", err)
		}
	}()

	return h
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.logger.Debug("Client registered", "id", client.Id)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				client.Active = false
				delete(h.clients, client)
				h.logger.Debug("Client unregistered", "id", client.Id)
			}
		case message := <-h.broadcast:
			if err := h.BroadcastMessage(message); err != nil {
				h.logger.Error("Error broadcasting message", "error", err)
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (h *Hub) BroadcastMessage(message Message) error {
	if len(h.clients) == 0 {
		h.logger.Debug("No clients to broadcast to")
		return nil
	}

	h.logger.Debug("Broadcasting message", "channel", message.Channel, "type", message.Type)

	j, err := message.ToJson()
	if err != nil {
		return err
	}

	for client := range h.clients {
		if message.From != nil && message.From.Id == client.Id {
			continue
		}
		if err := ws.Message.Send(client.conn, string(j)); err != nil {
			h.logger.Error("Error sending message to client", "id", client.Id, "error", err)
		}
	}

	return nil
}
