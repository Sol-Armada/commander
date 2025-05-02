package wsmanager

import "github.com/sol-armada/commander/operation"

func operationsChannelHandler(h *Hub, client *Client, message Message) error {
	switch message.Type {
	case MessageTypeCreate:
		op := operation.NewOperation(client.Id, message.Data["name"].(string))
		opMap, err := op.ToMap()
		if err != nil {
			return err
		}
		message.Channel = ChannelOperations
		message.Data = opMap
		message.From = client
		message.Type = MessageTypeCreated
		h.broadcast <- message
	case MessageTypeUpdate:
		op, err := operation.FromMap(message.Data)
		if err != nil {
			return err
		}

		updatedOp, err := operation.UpdateOperation(op, client.MemberId)
		if err != nil {
			return err
		}

		opMap, err := updatedOp.ToMap()
		if err != nil {
			return err
		}
		message.Data = opMap
		message.From = client
		message.Type = MessageTypeUpdated

		h.broadcast <- message
	}

	return nil
}

func (h *Hub) OperationsCreate(op *operation.Operation) error {
	j, err := op.ToMap()
	if err != nil {
		return err
	}
	message := Message{
		Type:    MessageTypeCreated,
		Data:    j,
		Channel: ChannelOperations,
	}
	h.broadcast <- message

	return nil
}
