package matchmaking

import (
	"errors"
	"iter"
)

var (
	ErrConcurrentUpdate = errors.New("concurrent update detected")
	ErrStoreIsFucked    = errors.New("store fucked up")
	ErrNoSuchMatch      = errors.New("match does not exist")
)

type Store interface {
	Fetch(id string) (*Match, error)
	Save(match *Match) error
	Delete(ids ...string)
	All() iter.Seq[*Matchstate]
}
