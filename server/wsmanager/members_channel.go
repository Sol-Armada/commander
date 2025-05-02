package wsmanager

import (
	"context"

	"github.com/sol-armada/commander/member"
)

func membersChannelHandler(h *Hub, client *Client, message Message) error {
	switch message.Type {
	case MessageTypeUpdate:
		h.logger.Debug("Updating member", "id", client.Id)
		
	}

	return nil
}

func (h *Hub) watchForMemberUpdates() error {
	h.logger.Debug("Watching for member updates")
	in := make(chan member.Member)
	if err := member.Watch(context.Background(), h.logger, in); err != nil {
		h.logger.Error("Failed to watch for member updates", "error", err)
		return err
	}

	for m := range in {
		h.logger.Debug("Received member update", "member", m)

		mMap, err := m.ToMap()
		if err != nil {
			h.logger.Error("Failed to convert member to map", "error", err)
			continue
		}

		message := Message{
			Type:    MessageTypeUpdated,
			Data:    mMap,
			Channel: ChannelMembers,
		}
		h.broadcast <- message
	}

	h.logger.Debug("Stopped watching for member updates")
	return nil
}
