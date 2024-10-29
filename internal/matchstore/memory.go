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
	id := match.Id

	entry, ok := m.matches.Load(id)

	if ok {
		if entry.Version != match.Version {
			return fmt.Errorf(
				"memorystore: attempted to save match '%s' with version '%d' which is different than last known version '%d'. %w",
				match.Id,
				match.Version,
				entry.Version,
				matchmaking.ErrConcurrentUpdate,
			)
		}
	}

	// We replace currentEntry we just got a moment ago with the new upToDateMatch
	// If "compare and swap" fails it means that the currentEntry we retrived has been modified in-between
	// If that is the case, we give up because a concurrent change happened before we could finish

	newEntry := match.Status()

	newEntry.Version = matchmaking.NextVersion() // Update the version before storing

	ok = m.matches.CompareAndSwap(id, entry, newEntry)

	if !ok {
		return fmt.Errorf(
			"memorystore: attempted to save match '%s' with new version '%d' but the match changed in-between",
			match.Id,
			match.Version,
			entry.Version,
			matchmaking.ErrConcurrentUpdate,
		)
	}

	return nil
}
