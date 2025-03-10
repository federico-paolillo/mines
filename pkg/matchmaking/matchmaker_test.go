package matchmaking_test

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestMatchmakerProducesNewMatchesAccordingToDifficulty(t *testing.T) {
	mm := matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			storage.NewInMemoryStore(),
		),
		generators.NewRngBoardGenerator(1234),
	)

	testCases := []struct {
		Difficulty game.Difficulty
		Width      int
		Height     int
		Lives      int
		State      game.Gamestate
	}{
		{
			Difficulty: game.BeginnerDifficulty,
			Width:      9,
			Height:     9,
			Lives:      2,
			State:      game.PlayingGame,
		},
		{
			Difficulty: game.IntermediateDifficulty,
			Width:      16,
			Height:     16,
			Lives:      1,
			State:      game.PlayingGame,
		},
		{
			Difficulty: game.ExpertDifficulty,
			Width:      30,
			Height:     16,
			Lives:      0,
			State:      game.PlayingGame,
		},
	}

	for _, testCase := range testCases {

		m, err := mm.New(
			123,
			testCase.Difficulty,
		)

		if err != nil {
			t.Fatalf(
				"matchmaker could not create match. %v",
				err,
			)
		}

		if m.Lives != testCase.Lives {
			t.Errorf(
				"lives generated for match with difficulty '%s' is wrong. expected %d got %d",
				testCase.Difficulty,
				m.Lives,
				testCase.Lives,
			)
		}

		if m.Width != testCase.Width {
			t.Errorf(
				"width generated for match with difficulty '%s' is wrong. expected %d got %d",
				testCase.Difficulty,
				m.Width,
				testCase.Width,
			)
		}

		if m.Height != testCase.Height {
			t.Errorf(
				"height generated for match with difficulty '%s' is wrong. expected %d got %d",
				testCase.Difficulty,
				m.Height,
				testCase.Height,
			)
		}

		if m.State != testCase.State {
			t.Errorf(
				"state generated for match with difficulty '%s' is wrong. expected '%s' got '%s'",
				testCase.Difficulty,
				m.State,
				testCase.State,
			)
		}
	}
}

func TestMatchmakerWillNotProcessMovesForNonExistingMatches(t *testing.T) {
	mm := matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			storage.NewInMemoryStore(),
		),
		generators.NewRngBoardGenerator(1234),
	)
	_, err := mm.Apply(
		"does-not-exist",
		matchmaking.Move{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		},
	)

	if !errors.Is(err, matchmaking.ErrNoSuchMatch) {
		t.Errorf(
			"matchmaker did not complain that the match does not exist. %v",
			err,
		)
	}
}

func TestMatchmakerWillProcessMoves(t *testing.T) {
	mm := matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			storage.NewInMemoryStore(),
		),
		generators.NewRngBoardGenerator(1234),
	)

	m, _ := mm.New(
		123,
		game.BeginnerDifficulty,
	)

	_, err := mm.Apply(
		m.Id,
		matchmaking.Move{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		},
	)

	if err != nil {
		t.Fatalf(
			"matchmaker failed to apply moves. %v",
			err,
		)
	}
}

func TestMatchmakerReportsConcurrencyCollision(t *testing.T) {
	mm := matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			storage.NewInMemoryStore(),
		),
		generators.NewRngBoardGenerator(1234),
	)

	m, _ := mm.New(
		123,
		game.BeginnerDifficulty,
	)

	randomSleep := func() {
		jitter := rand.Intn(10)
		sleepTime := 10 + jitter
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	errStream := make(chan error, 3)

	var wg sync.WaitGroup
	var err error

	moveSpammer := func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				randomSleep()

				_, err := mm.Apply(
					m.Id,
					matchmaking.Move{
						Type: matchmaking.MoveOpen,
						X:    1,
						Y:    1,
					},
				)

				if errors.Is(err, matchmaking.ErrConcurrentUpdate) {
					select {
					case errStream <- err:
					default:
					}
				}
			}
		}
	}

	spammersCount := 3
	for i := 0; i < spammersCount; i++ {
		wg.Add(1)
		go moveSpammer(ctx)
	}

	select {
	case <-ctx.Done():
		t.Fatal("timed-out without reporting concurrency issues")
	case err = <-errStream:
		cancel()
	}

	wg.Wait()

	if !errors.Is(err, matchmaking.ErrConcurrentUpdate) {
		t.Fatalf(
			"did not report concurrency issues. %v",
			err,
		)
	}
}

func TestMatchmakerWillPersistMoves(t *testing.T) {
	mm := matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			storage.NewInMemoryStore(),
		),
		generators.NewRngBoardGenerator(1234),
	)

	m, _ := mm.New(
		123,
		game.BeginnerDifficulty,
	)

	_, err := mm.Apply(
		m.Id,
		matchmaking.Move{
			Type: matchmaking.MoveFlag,
			X:    1,
			Y:    1,
		},
	)

	if err != nil {
		t.Fatalf(
			"matchmake failed to apply moves. %v",
			err,
		)
	}

	state, err := mm.Get(m.Id)

	if err != nil {
		t.Fatalf(
			"matchmaker failed to get match. %v",
			err,
		)
	}

	if state.Cells[0][0].State != board.FlaggedCell {
		t.Log("matchmaker did not persist move")
	}
}
