package storage

import (
	"errors"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

var (
	ErrNoSuchItem       = errors.New("store: item not found")
	ErrConcurrentUpdate = errors.New("store: concurrent update detected")
)

type Store interface {
	Fetch(id string) (*matchmaking.Matchstate, error)
	Save(matchstate *matchmaking.Matchstate) error
	Healthy() error
}
