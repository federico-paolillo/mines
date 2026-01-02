package testutils

import (
	"github.com/federico-paolillo/mines/internal/id"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

const SomeMatchId = "abc"

const SomeMatchVersion = 123

func SomeMatch() *matchmaking.Match {
	return SomeCustomMatch("abc", 123, 456)
}

func SomeRandomMatch() *matchmaking.Match {
	return SomeCustomMatch(
		id.Generate(),
		123,
		456,
	)
}

func SomeCustomMatch(
	id string,
	version uint64,
	startTime int64,
) *matchmaking.Match {
	bb := board.NewBuilder(
		dimensions.Size{
			Width:  2,
			Height: 1,
		},
	)

	_ = bb.PlaceSafe(1, 1)
	_ = bb.PlaceMine(2, 1)

	b := bb.Build()

	g := game.NewGame(12, b)

	m := matchmaking.NewMatch(
		id,
		version,
		startTime,
		b,
		g,
	)

	return m
}
