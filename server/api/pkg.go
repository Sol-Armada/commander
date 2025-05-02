package api

import (
	"log/slog"

	"github.com/sol-armada/commander/wsmanager"
	"github.com/sol-armada/sol-bot/stores"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	db *stores.Client
	wm *wsmanager.Hub
}

type ApiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data"`
}

func NewServer(logger *slog.Logger, db *stores.Client) Server {
	wm := wsmanager.NewHub(logger, db)
	go wm.Run()
	return Server{
		db: db,
		wm: wm,
	}
}

func (s *Server) DB() *stores.Client {
	return s.db
}
