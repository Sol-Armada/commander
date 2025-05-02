package squad

import "github.com/sol-armada/commander/member"

type Squad struct {
	Id      string           `json:"id"`
	Name    string           `json:"name"`
	Members []*member.Member `json:"members"`
}
