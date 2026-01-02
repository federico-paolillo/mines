package storage

import (
	"iter"

	"github.com/federico-paolillo/mines/internal/reaper"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type ReaperStore struct {
	memstore Store
}

func NewReaperStore(store Store) *ReaperStore {
	return &ReaperStore{
		store,
	}
}

func (g *ReaperStore) All() iter.Seq[*matchmaking.Matchstate] {
	return g.memstore.All()
}

func (g *ReaperStore) Delete(ids ...string) {
	g.memstore.Delete(ids...)
}

var _ reaper.Store = (*ReaperStore)(nil)
