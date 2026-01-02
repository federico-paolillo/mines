package storage

import (
	"errors"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type MatchStore struct {
	memstore Store
}

func NewMatchStore(store Store) *MatchStore {
	return &MatchStore{
		store,
	}
}

func (m *MatchStore) Fetch(id string) (*matchmaking.Match, error) {
	matchstate, err := m.memstore.Fetch(id)

	if errors.Is(err, ErrNoSuchItem) {
		return nil, fmt.Errorf(
			"matchstore: could not find match with id '%s'. %w",
			id,
			matchmaking.ErrNoSuchMatch,
		)
	}

	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak internal error types from this interface
		return nil, fmt.Errorf(
			"matchstore: failed to find match with id '%s'. %w. %v",
			id,
			matchmaking.ErrStoreIsFucked,
			err,
		)
	}

	match := HydrateMatch(matchstate)

	return match, nil
}

func (m *MatchStore) Save(match *matchmaking.Match) error {
	matchstate := match.Status()

	err := m.memstore.Save(matchstate)

	if errors.Is(err, ErrConcurrentUpdate) {
		return fmt.Errorf(
			"matchstore: could not save match with id '%s'. %w",
			match.Id,
			matchmaking.ErrConcurrentUpdate,
		)
	}

	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak internal error types from this interface
		return fmt.Errorf(
			"matchstore: failed to save match with id '%s'. %w. %v",
			match.Id,
			matchmaking.ErrStoreIsFucked,
			err,
		)
	}

	return nil
}

var _ matchmaking.Store = (*MatchStore)(nil)
