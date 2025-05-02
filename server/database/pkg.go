package database

import (
	"context"
	"errors"

	"github.com/sol-armada/sol-bot/stores"
)

type Options struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Database       string
	ReplicaSetName string
}

func New(ctx context.Context, opts *Options) (*stores.Client, error) {
	c, err := stores.New(ctx, opts.Host, opts.Port, opts.Username, opts.Password, opts.Database, opts.ReplicaSetName)
	if err != nil {
		return nil, err
	}

	if !c.Connected() {
		return nil, errors.New("failed to connect to database")
	}

	return c, nil
}
