package db

import (
	"time"

	"github.com/ATtendev/share/config"
	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/store/db/ent"
)

type Store struct {
	ent     *ent.Client
	timeOut time.Duration
}

func NewConnection(cfg *config.DatabaseConf) *Store {
	db := ent.NewClient(
		ent.Log(log.Info), // logger
		ent.Driver(cfg.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	log.Info("databse connected")
	return &Store{
		ent:     db,
		timeOut: 3000,
	}
}

func (s *Store) Close() error {
	return s.ent.Close()
}
