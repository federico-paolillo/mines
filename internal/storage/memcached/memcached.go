package memcached

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type Memcached struct {
	client *memcache.Client
	ttl    time.Duration
}

func NewMemcached(client *memcache.Client, ttl time.Duration) *Memcached {
	return &Memcached{
		client,
		ttl,
	}
}

func (m *Memcached) Fetch(id string) (*matchmaking.Matchstate, error) {
	item, err := m.client.Get(id)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil, fmt.Errorf("memcached: could not find matchstate '%s': %w", id, storage.ErrNoSuchItem)
	}

	if err != nil {
		return nil, fmt.Errorf("memcached: failed to fetch matchstate '%s': %w", id, err)
	}

	var matchstateJSON MatchstateJSON

	err = json.Unmarshal(item.Value, &matchstateJSON)
	if err != nil {
		return nil, fmt.Errorf("memcached: failed to unmarshal matchstate '%s': %w", id, err)
	}

	matchstate := JSONToMatchstate(&matchstateJSON)
	matchstate.Version = item.CasID

	return matchstate, nil
}

func (m *Memcached) Save(matchstate *matchmaking.Matchstate) error {
	matchstateJSON := MatchstateToJSON(matchstate)

	value, err := json.Marshal(matchstateJSON)
	if err != nil {
		return fmt.Errorf("memcached: failed to marshal matchstate '%s': %w", matchstate.Id, err)
	}

	item := &memcache.Item{
		Key:        matchstate.Id,
		Value:      value,
		Expiration: int32(m.ttl.Seconds()),
	}

	if matchstate.Version == 0 {
		err = m.client.Add(item)
	} else {
		item.CasID = matchstate.Version
		err = m.client.CompareAndSwap(item)
	}

	if errors.Is(err, memcache.ErrCASConflict) {
		return fmt.Errorf("memcached: concurrent update on matchstate '%s': %w", matchstate.Id, storage.ErrConcurrentUpdate)
	}

	if err != nil {
		return fmt.Errorf("memcached: failed to save matchstate '%s': %w", matchstate.Id, err)
	}

	// After a successful save, we need to update the matchstate with the new version (CasID).
	// This is essential for subsequent updates to succeed.
	// Unfortunately, memcache's CompareAndSwap doesn't return the new CasID.
	// We must fetch the item again to get it. This introduces a small race condition window:
	// another process could update the item between our Save and Get, causing us to retrieve
	// a CasID that is already outdated. For this application's purpose this is an accepted tradeoff
	// for simplicity, as the alternative (e.g., distributed locks) is too complex.
	newItem, err := m.client.Get(matchstate.Id)
	if err != nil {
		return fmt.Errorf("memcached: failed to update casid for '%s': %w", matchstate.Id, err)
	}

	matchstate.Version = newItem.CasID

	return nil
}

func (m *Memcached) Healthy() error {
	err := m.client.Ping()

	return fmt.Errorf(
		"memcached: failed to ping: %w",
		err,
	)
}

var _ storage.Store = (*Memcached)(nil)
