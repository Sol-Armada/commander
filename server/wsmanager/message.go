package wsmanager

import (
	"encoding/json"
	"errors"
)

type Type string

const (
	MessageTypeRegistered Type = "registered"
	MessageTypePing       Type = "ping"
	MessageTypeList       Type = "list"
	MessageTypeUpdated    Type = "updated"
	MessageTypeRemoved    Type = "removed"
	MessageTypeCreated    Type = "created"
	MessageTypeDeleted    Type = "deleted"
	MessageTypeCreate     Type = "create"
	MessageTypeDelete     Type = "delete"
	MessageTypeUpdate     Type = "update"
)

type Channel string

const (
	ChannelRegister   Channel = "register"
	ChannelOperations Channel = "operations"
	ChannelMembers    Channel = "members"
	ChannelKeepAlive  Channel = "keepalive"
)

type Message struct {
	Channel Channel        `json:"channel"`
	Type    Type           `json:"type"`
	Data    map[string]any `json:"data"`
	From    *Client        `json:"-"`
}

func (m *Message) ToJson() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to marshal response"))
	}
	return b, nil
}
