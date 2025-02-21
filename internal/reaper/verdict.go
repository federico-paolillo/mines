package reaper

import (
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type ReapingVerdict = string

var (
	Ok        ReapingVerdict = "ok"
	Expired                  = "expired"
	Completed                = "completed"
)

func emitVerdict(m *matchmaking.Matchstate) ReapingVerdict {
	if isExpired(m) {
		return Expired
	}

	if isCompleted(m) {
		return Completed
	}

	return Ok
}

func isExpired(_ *matchmaking.Matchstate) bool {
	return false // TODO: Missing timestamp on Matchstate
}

func isCompleted(m *matchmaking.Matchstate) bool {
	return m.State != game.PlayingGame
}

func verdictIsUnfavourable(v ReapingVerdict) bool {
	return v == Expired || v == Completed
}
