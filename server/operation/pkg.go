package operation

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/rs/xid"
	"github.com/sol-armada/commander/member"
	"github.com/sol-armada/commander/ship"
	"github.com/sol-armada/commander/squad"
	"github.com/sol-armada/commander/utils"
)

type OperationMember struct {
	Connected bool `json:"connected"`

	member.Member
}

type Operation struct {
	Id        string             `json:"id"`
	Name      string             `json:"name"`
	OnStandBy []*OperationMember `json:"onStandBy"`
	Ships     []*ship.Ship       `json:"ships"`
	Squads    []*squad.Squad     `json:"squads"`

	Creator       string    `json:"creator"`
	Created       time.Time `json:"created"`
	LastUpdatedBy string    `json:"lastUpdatedBy"`
	Updated       time.Time `json:"updated"`

	mu sync.Mutex `json:"-"`
}

var mu sync.Mutex
var operations = map[string]*Operation{}

func NewOperation(creator string, name string) *Operation {
	if name == "" {
		name = utils.RandomName()
	}

	op := &Operation{
		Id:        xid.New().String(),
		Name:      name,
		OnStandBy: []*OperationMember{},
		Ships:     []*ship.Ship{},
		Squads:    []*squad.Squad{},
		Creator:   creator,
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	mu.Lock()
	defer mu.Unlock()
	operations[op.Id] = op

	return op
}

func GetOperations() []*Operation {
	mu.Lock()
	defer mu.Unlock()

	ops := []*Operation{}
	for _, op := range operations {
		ops = append(ops, op)
	}
	ops = removeDuplicates(ops)
	return ops
}

func GetOperation(id string) (*Operation, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, op := range operations {
		if op.Id == id {
			return op, nil
		}
	}

	return nil, utils.ErrNotFound
}

func UpdateOperation(op *Operation, updatedBy string) (*Operation, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := operations[op.Id]; ok {
		op.mu.Lock()
		defer op.mu.Unlock()

		op.LastUpdatedBy = updatedBy
		op.Updated = time.Now()
		operations[op.Id] = op
		return op, nil
	}

	return nil, utils.ErrNotFound
}

func removeDuplicates(ops []*Operation) []*Operation {
	keys := make(map[string]bool)
	list := []*Operation{}

	for _, op := range ops {
		if _, value := keys[op.Id]; !value {
			keys[op.Id] = true
			list = append(list, op)
		}
	}

	return list
}

func FromMap(m map[string]any) (*Operation, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var op *Operation
	if err := json.Unmarshal(j, &op); err != nil {
		return nil, err
	}

	return op, nil
}

func (op *Operation) ToMap() (map[string]any, error) {
	b, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]any
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return nil, err
	}
	return jsonMap, nil
}
