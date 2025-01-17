package storage

import (
	"iter"

	"github.com/federico-paolillo/mines/internal/gc"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type GcStore struct {
	memstore *InMemoryStore
}

func (g *GcStore) All() iter.Seq[*matchmaking.Matchstate] {
	return g.memstore.All()
}

func (g *GcStore) Delete(ids ...string) {
	g.memstore.Delete(ids...)
}

var _ gc.Store = (*GcStore)(nil)
