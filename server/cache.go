package main

import (
	"encoding/json"
	"log/slog"

	valkeyapi "github.com/valkey-io/valkey-glide/go/api"
	"github.com/valkey-io/valkey-glide/go/api/options"
)

type cache struct {
	valkeyapi.GlideClientCommands
}

func (c *cache) Set(key string, value any) error {
	valBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.GlideClientCommands.Set(key, string(valBytes))
	return err
}

func (c *cache) SetWithTTL(key string, value any, ttl int) error {
	valBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var opts options.SetOptions
	if ttl > 0 {
		opts = options.SetOptions{
			Expiry: &options.Expiry{
				Type:  options.Seconds,
				Count: uint64(ttl),
			},
		}
	}
	_, err = c.GlideClientCommands.SetWithOptions(key, string(valBytes), opts)
	return err
}

func (c *cache) Get(key string) (any, error) {
	res, err := c.GlideClientCommands.Get(key)
	if err != nil {
		return nil, err
	}

	if res.IsNil() {
		return nil, nil
	}

	var result any
	if err := json.Unmarshal([]byte(res.Value()), &result); err != nil {
		slog.Error("failed to unmarshal result", "err", err)
		return nil, err
	}

	return result, nil
}

func (c *cache) Delete(key string) error {
	_, err := c.GlideClientCommands.Del([]string{key})
	return err
}
