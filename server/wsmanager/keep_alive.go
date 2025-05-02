package wsmanager

import (
	"time"

	"github.com/sol-armada/commander/member"
)

func (h *Hub) KeepAlive() {
	for {
		if len(h.clients) >= 0 {
			h.broadcast <- Message{
				Channel: ChannelKeepAlive,
				Type:    MessageTypePing,
				Data:    map[string]any{},
				From:    nil,
			}
		}

		time.Sleep(time.Second * 15)

		for client := range h.clients {
			if time.Since(client.LastSeen) > time.Second*30 {
				h.unregister <- client
			}
		}
	}
}

func keepAliveChannelHandler(h *Hub, client *Client, message Message) error {
	client.LastSeen = time.Now()
	if message.Type != MessageTypePing {
		return nil
	}

	client.Active = true

	m, err := member.Get(h.db, client.MemberId)
	if err != nil {
		return err
	}

	m.Active = true

	mMap, err := m.ToMap()
	if err != nil {
		return err
	}

	message = Message{
		Channel: ChannelMembers,
		Type:    MessageTypeUpdated,
		Data:    mMap,
	}

	message.From = client
	select {
	case h.broadcast <- message:
	case <-time.After(time.Second * 5):
		// Timeout to avoid blocking the hub
		// This is a simple way to avoid blocking the hub
		// but you can also use a buffered channel
		// or a separate goroutine to handle the message
		h.logger.Warn("Timeout sending message to hub", "message", message)
	}

	return nil
}
