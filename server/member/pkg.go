package member

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"sort"

	solmembers "github.com/sol-armada/sol-bot/members"
	"github.com/sol-armada/sol-bot/ranks"
	"github.com/sol-armada/sol-bot/stores"
)

type Member struct {
	Id     string     `json:"id" bson:"_id"`
	Name   string     `json:"name"`
	Rank   ranks.Rank `json:"rank"`
	Active bool       `json:"active"`
}

func fromSolMember(member solmembers.Member) Member {
	return Member{
		Id:     member.Id,
		Name:   member.Name,
		Rank:   member.Rank,
		Active: false,
	}
}

func Get(dbc *stores.Client, id string) (*Member, error) {
	ctx := context.TODO()
	s, ok := dbc.GetMembersStore()
	if !ok {
		return nil, errors.New("failed to get members store")
	}

	cur, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if !cur.Next(ctx) {
		return nil, errors.New("member not found")
	}

	member := &Member{}
	if err := cur.Decode(member); err != nil {
		return nil, err
	}

	return member, nil
}

func List(dbc *stores.Client) []*Member {
	membersRaw, err := solmembers.List(0)
	if err != nil {
		slog.Error("failed to get members", "err", err)
		os.Exit(1)
	}

	members := make(map[string]*Member)
	for _, member := range membersRaw {
		if member.IsBot {
			continue
		}

		if member.Rank == ranks.None {
			member.Rank = ranks.Guest
		}

		m := &Member{
			Id:   member.Id,
			Name: member.Name,
			Rank: member.Rank,
		}

		members[member.Id] = m
	}

	return toSlice(sortMembers(members))
}

func toSlice(members map[string]*Member) []*Member {
	slice := make([]*Member, 0, len(members))
	for _, member := range members {
		slice = append(slice, member)
	}
	return slice
}

func sortMembers(members map[string]*Member) map[string]*Member {
	keys := make([]string, 0, len(members))
	for k := range members {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if members[keys[i]].Rank > members[keys[j]].Rank {
			return false
		}
		if members[keys[i]].Rank < members[keys[j]].Rank {
			return true
		}
		if members[keys[i]].Name < members[keys[j]].Name {
			return true
		}
		if members[keys[i]].Name > members[keys[j]].Name {
			return false
		}
		return false
	})

	nMembers := make(map[string]*Member)
	for _, k := range keys {
		nMembers[k] = members[k]
	}

	return nMembers
}

func (m *Member) ToMap() (map[string]any, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	if err := json.Unmarshal(j, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func Watch(ctx context.Context, logger *slog.Logger, out chan Member) error {
	membersChan := make(chan solmembers.Member)

	go func() {
		if err := solmembers.Watch(ctx, membersChan); err != nil {
			logger.Error("failed to watch members", "err", err)
			return
		}
	}()

	go func(c chan solmembers.Member) {
		logger.Debug("watching for member updates")
		for {
			select {
			case <-ctx.Done():
				return
			case solmember := <-c:
				if solmember.IsBot {
					continue
				}

				logger.Debug("received member update", "member", solmember.Id)

				out <- fromSolMember(solmember)
			}
		}
	}(membersChan)
	return nil
}
