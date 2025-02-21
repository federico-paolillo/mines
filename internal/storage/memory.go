package storage

import (
	"errors"
	"fmt"
	"iter"
	"sync"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MatchstatesMap = map[string]*matchmaking.Matchstate

var (
	ErrNoSuchItem       = errors.New("memstore: item not found")
	ErrConcurrentUpdate = errors.New("memstore: concurrent update detected")
)

type InMemoryStore struct {
	mu     sync.RWMutex
	states MatchstatesMap
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		states: make(MatchstatesMap),
	}
}

func (m *InMemoryStore) Fetch(id string) (*matchmaking.Matchstate, error) {
	m.mu.RLock()

	// Possible improvement: unlock before hydrating so that we don't keep the lock for nothing
	defer m.mu.RUnlock()

	if entry, ok := m.states[id]; ok {
		return entry, nil
	}

	return nil, fmt.Errorf(
		"memstore: could not find matchstate '%s'. %w",
		id,
		ErrNoSuchItem,
	)
}

func (m *InMemoryStore) Save(matchstate *matchmaking.Matchstate) error {
	m.mu.Lock()

	defer m.mu.Unlock()

	newEntry := cloneMatchstate(matchstate)

	// You can update a Matchstate only if the version you provide is still the same as what's in store
	// This optimistic concurrency token will ensure that we do not overwrite newer versions

	if existingEntry, ok := m.states[newEntry.Id]; ok {
		if existingEntry.Version != newEntry.Version {
			return fmt.Errorf(
				"memstore: attempted to save matchstate '%s' with version '%d' which is different than last known version '%d'. %w",
				newEntry.Id,
				newEntry.Version,
				existingEntry.Version,
				ErrConcurrentUpdate,
			)
		}
	}

	// Change the version before storing

	newEntry.Version = matchmaking.NextVersion()

	m.states[newEntry.Id] = newEntry

	return nil
}

func (m *InMemoryStore) All() iter.Seq[*matchmaking.Matchstate] {
	return func(yield func(s *matchmaking.Matchstate) bool) {
		m.mu.RLock()

		defer m.mu.RUnlock()

		// TODO: What if yield blocks indefinitely and RLock never gets released ?

		for _, v := range m.states {
			if !yield(v) {
				break
			}
		}
	}
}

func (m *InMemoryStore) Delete(ids ...string) {
	m.mu.Lock()

	defer m.mu.Unlock()

	for _, v := range ids {
		delete(m.states, v)
	}
}

func cloneMatchstate(matchstate *matchmaking.Matchstate) *matchmaking.Matchstate {
	cells := make(matchmaking.Cells, 0, len(matchstate.Cells))

	for _, row := range matchstate.Cells {
		cols := make([]matchmaking.Cell, 0, len(row))

		for _, cell := range row {
			cols = append(
				cols,
				matchmaking.Cell{
					X:     cell.X,
					Y:     cell.Y,
					Mined: cell.Mined,
					State: cell.State,
				},
			)
		}

		cells = append(cells, cols)
	}

	return &matchmaking.Matchstate{
		Id:      matchstate.Id,
		Version: matchstate.Version,
		Lives:   matchstate.Lives,
		State:   matchstate.State,
		Width:   matchstate.Width,
		Height:  matchstate.Height,
		Cells:   cells,
	}
}
