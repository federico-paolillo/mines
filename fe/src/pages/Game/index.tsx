import { useRoute } from "preact-iso";
import { Countdown } from "../../components/Countdown";
import { LifeContainer } from "../../components/LifeContainer";
import { MinesweeperBoard } from "../../components/Minesweeper";
import { Spinner } from "../../components/Spinner";
import { useGameState } from "../../hooks/useGameState";

export function Game() {
  const { query } = useRoute();
  const params = new URLSearchParams(query);
  const gameId = params.get("id");

  const { gameState, loading, onOpenCell, onFlagCell, onExpired } =
    useGameState(gameId);

  if (loading) {
    return <Spinner isOpen={true} />;
  }

  if (gameState) {
    return (
      <div class="flex flex-col items-center justify-center min-h-screen bg-[#008080]">
        <div class="mb-4 text-white text-xl font-bold flex flex-col items-center gap-2">
          <span>{`Game: ${gameId}`}</span>
          <Countdown
            startTime={gameState.startTime ?? 0}
            durationSeconds={2 * 60 * 60}
            onExpired={onExpired}
          />
          <LifeContainer lives={gameState.lives ?? 0} />
        </div>
        <MinesweeperBoard
          gameState={gameState}
          onOpenCell={onOpenCell}
          onFlagCell={onFlagCell}
        />
      </div>
    );
  }

  return <div class="text-white">Game not found or failed to load.</div>;
}
