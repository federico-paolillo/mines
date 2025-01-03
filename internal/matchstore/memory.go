package matchstore

import (
	"fmt"
	"sync"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MatchesMap = map[string]*matchmaking.Matchstate

type MemoryStore struct {
	mu      sync.RWMutex
	matches MatchesMap
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		matches: make(MatchesMap),
	}
}

var _ matchmaking.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Fetch(id string) (*matchmaking.Match, error) {
	m.mu.RLock()

	// Possible improvement: unlock before hydrating so that we don't keep the lock for nothing
	defer m.mu.RUnlock()

	if entry, ok := m.matches[id]; ok {
		match := HydrateMatch(entry)

		return match, nil
	}

	return nil, fmt.Errorf(
		"memorystore: could not find match '%s'. %w",
		id,
		matchmaking.ErrNoSuchMatch,
	)
}

func (m *MemoryStore) Save(match *matchmaking.Match) error {
	m.mu.Lock()

	defer m.mu.Unlock()

	newEntry := match.Status()

	// You can update a Match only if the version you provide is still the same as what's in store
	// This optimistic concurrency token will ensure that we do not overwrite newer versions

	if existingEntry, ok := m.matches[newEntry.Id]; ok {
		if existingEntry.Version != newEntry.Version {
			return fmt.Errorf(
				"memorystore: attempted to save match '%s' with version '%d' which is different than last known version '%d'. %w",
				newEntry.Id,
				newEntry.Version,
				existingEntry.Version,
				matchmaking.ErrConcurrentUpdate,
			)
		}
	}

	// Change the version before storing
	// This entry has not yet left the store so it is safe to change in place

	newEntry.Version = matchmaking.NextVersion()

	m.matches[newEntry.Id] = newEntry

	return nil
}
