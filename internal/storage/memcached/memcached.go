package memcached

import (
	"iter"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type Memcached struct {
	client *memcache.Client
}

func NewMemcached(client *memcache.Client) *Memcached {
	return &Memcached{
		client,
	}
}

func (m *Memcached) Fetch(id string) (*matchmaking.Matchstate, error) {
	// if errors.Is(err, memcache.ErrCacheMiss) {
	// 	return nil, fmt.Errorf(
	// 		"memcached: could not find matchstate '%s'. %w",
	// 		id,
	// 		matchmaking.ErrNoSuchMatch,
	// 	)
	// }
	// if err != nil {
	// 	//nolint:errorlint // We do not want to wrap and leak internal error types from this interface
	// 	return nil, fmt.Errorf(
	// 		"memcached: failed to fetch matchstate '%s'. %v. %w",
	// 		id,
	// 		err,
	// 		matchmaking.ErrStoreIsFucked,
	// 	)
	// }
	// var matchstate *matchmaking.Matchstate
	// err = json.Unmarshal(item.Value, &matchstate)
	// if err != nil {
	// 	//nolint:errorlint // We do not want to wrap and leak internal error types from this interface
	// 	return nil, fmt.Errorf(
	// 		"memcached: failed to unmarshal matchstate '%s'. %v. %w",
	// 		id,
	// 		err,
	// 		matchmaking.ErrStoreIsFucked,
	// 	)
	// }
	// matchstate.Version = item.CasID
	// return matchstate, nil
	// item, err := m.client.Get(id)
	panic("implement me")
}

func (m *Memcached) Save(matchstate *matchmaking.Matchstate) error {
	panic("implement me")
}

func (m *Memcached) All() iter.Seq[*matchmaking.Matchstate] {
	panic("implement me")
}

func (m *Memcached) Delete(ids ...string) {
	panic("implement me")
}

var _ storage.Store = (*Memcached)(nil)
