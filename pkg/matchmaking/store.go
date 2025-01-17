package matchmaking

import (
	"errors"
)

var (
	ErrConcurrentUpdate = errors.New("concurrent update detected")
	ErrNoSuchMatch      = errors.New("match does not exist")
	ErrStoreIsFucked    = errors.New("store is broken")
)

type Store interface {
	Fetch(id string) (*Match, error)
	Save(match *Match) error
}
