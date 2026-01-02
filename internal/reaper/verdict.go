package reaper

import (
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type ReapingVerdict = string

const TwoHoursInSeconds = 7200

var (
	Ok        ReapingVerdict = "ok"
	Expired                  = "expired"
	Completed                = "completed"
)

func emitVerdict(
	now int64,
	m *matchmaking.Matchstate,
) ReapingVerdict {
	if isExpired(now, m) {
		return Expired
	}

	if isCompleted(m) {
		return Completed
	}

	return Ok
}

func isExpired(now int64, m *matchmaking.Matchstate) bool {
	return (now - m.StartTime) > TwoHoursInSeconds
}

func isCompleted(m *matchmaking.Matchstate) bool {
	return m.State != game.PlayingGame
}

func verdictIsUnfavourable(v ReapingVerdict) bool {
	return v == Expired || v == Completed
}
