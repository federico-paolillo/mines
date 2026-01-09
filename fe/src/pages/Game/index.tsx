import { useEffect, useState } from "preact/hooks";
import { useLocation, useRoute } from "preact-iso";
import type { MatchstateDto } from "../../client/models/res";
import { useApiClient } from "../../clientContext";
import { Countdown } from "../../components/Countdown";
import { MinesweeperBoard } from "../../components/Minesweeper";
import { Spinner } from "../../components/Spinner";

export function Game() {
  const { route } = useLocation();
  const { query } = useRoute();
  const params = new URLSearchParams(query);
  const gameId = params.get("id");
  const client = useApiClient();

  const [gameState, setGameState] = useState<MatchstateDto | null>(null);
  const [loading, setLoading] = useState(!!gameId);

  useEffect(() => {
    if (!gameId) {
      setLoading(false);
      return;
    }

    const fetchGame = async () => {
      setLoading(true);
      const result = await client.fetchMatch(gameId);
      setLoading(false);

      if (result.success) {
        setGameState(result.value);
      } else {
        console.error(result.error);
      }
    };

    fetchGame();
  }, [gameId, client]);

  const handleCellClick = (x: number, y: number) => {
    console.log(`Clicked cell at ${x}, ${y}`);
  };

  const handleCellRightClick = (x: number, y: number) => {
    console.log(`Right clicked cell at ${x}, ${y}`);
  };

  const handleExpired = () => {
    route("/game-over");
  };

  return (
    <div class="flex flex-col items-center justify-center min-h-screen bg-[#008080]">
      <Spinner isOpen={loading} />
      <div class="mb-4 text-white text-xl font-bold flex flex-col items-center gap-2">
        <span>{gameId ? `Game: ${gameId}` : "Minesweeper Demo"}</span>
        {gameState?.startTime && (
          <Countdown
            startTime={gameState.startTime}
            onExpired={handleExpired}
          />
        )}
      </div>
      {gameState ? (
        <MinesweeperBoard
          gameState={gameState}
          onCellClick={handleCellClick}
          onCellRightClick={handleCellRightClick}
        />
      ) : (
        !loading && (
          <div class="text-white">Game not found or failed to load.</div>
        )
      )}
    </div>
  );
}
