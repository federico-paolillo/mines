package gc

import (
	"iter"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type Store interface {
	Delete(ids ...string)
	All() iter.Seq[*matchmaking.Matchstate]
}
