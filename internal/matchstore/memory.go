package matchstore

import (
	"fmt"

	"github.com/federico-paolillo/mines/internal/syncmapt"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MemoryStore struct {
	matches syncmapt.SyncMap[string, *matchmaking.Matchstate]
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

var _ matchmaking.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Fetch(id string) (*matchmaking.Match, error) {
	entry, ok := m.matches.Load(id)

	if !ok {
		return nil, fmt.Errorf(
			"memorystore: could not find match '%s'. %w",
			id,
			matchmaking.ErrNoSuchMatch,
		)
	}

	match := HydrateMatch(entry)

	return match, nil
}

func (m *MemoryStore) Save(match *matchmaking.Match) error {
	newEntry := match.Status()

	currentEntry, exists := m.matches.Load(newEntry.Id)

	if exists {
		return m.optimisticSwap(currentEntry, newEntry)
	}

	// Change the version before storage
	// This entry has not yet left the storage so it is safe to change in place

	newEntry.Version = matchmaking.NextVersion()

	m.matches.Store(newEntry.Id, newEntry)

	return nil
}

func (m *MemoryStore) optimisticSwap(
	currentEntry *matchmaking.Matchstate,
	newEntry *matchmaking.Matchstate,
) error {
	id := currentEntry.Id

	if currentEntry.Version != newEntry.Version {
		return fmt.Errorf(
			"memorystore: attempted to save match '%s' with version '%d' which is different than last known version '%d'. %w",
			id,
			newEntry.Version,
			currentEntry.Version,
			matchmaking.ErrConcurrentUpdate,
		)
	}

	// We replace currentEntry we just got a moment ago with the new upToDateMatch
	// If "compare and swap" fails it means that the currentEntry we retrieved has been modified in-between
	// If that is the case, we give up because a concurrent change happened before we could finish

	// Change the version before storage
	// This entry has not yet left the storage so it is safe to change in place
	newEntry.Version = matchmaking.NextVersion()

	// Compare and swap checks that currentEntry is equal to the one in the Map before replacing it with newEntry

	didSwap := m.matches.CompareAndSwap(id, currentEntry, newEntry)

	if !didSwap {
		return fmt.Errorf(
			"memorystore: attempted to save match '%s' with new version '%d' but the match changed in-between. %w",
			id,
			newEntry.Version,
			matchmaking.ErrConcurrentUpdate,
		)
	}

	return nil
}
