package position

import "github.com/sol-armada/commander/member"

type Position struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Member *member.Member `json:"member"`
}
