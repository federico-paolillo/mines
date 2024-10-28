package matchmaking

import "errors"

var ErrConcurrentUpdate = errors.New("concurrent update detected")
var ErrStoreIsFucked = errors.New("store fucked up")
var ErrNoSuchMatch = errors.New("match does not exist")

type Store interface {
	Fetch(id string) (*Match, error)
	Save(id string) error
}
