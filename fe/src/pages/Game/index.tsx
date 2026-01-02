import { useLocation, useRoute } from "preact-iso";
import { useState, useEffect } from "preact/hooks";
import { MinesweeperBoard } from "../../components/Minesweeper";
import { MatchstateDto, CellDto } from "../../client/models/res";
import { CellstateObject } from "../../client/models/board";
import { GamestateObject } from "../../client/models/game";

export function Game() {
  const { query } = useRoute();
  const params = new URLSearchParams(query);
  const gameId = params.get("id");

  // Mock initial state for demonstration
  // In a real app, this would be fetched from the backend using the gameId
  const [gameState, setGameState] = useState<MatchstateDto | null>(null);

  useEffect(() => {
    // Simulate fetching data
    // Mock data
    const width = 9;
    const height = 9;
    const cells: CellDto[] = [];
    for (let y = 1; y <= height; y++) {
      for (let x = 1; x <= width; x++) {
        cells.push({
          state: CellstateObject.Closed,
          x,
          y,
        });
      }
    }

    setGameState({
      id: gameId ?? "demo",
      lives: 3,
      state: GamestateObject.Playing,
      width,
      height,
      cells,
    });
  }, [gameId]);

  const handleCellClick = (x: number, y: number) => {
    console.log(`Clicked cell at ${x}, ${y}`);
  };

  const handleCellRightClick = (x: number, y: number) => {
    console.log(`Right clicked cell at ${x}, ${y}`);
  };

  return (
    <div class="flex flex-col items-center justify-center min-h-screen bg-[#008080]">
      <div class="mb-4 text-white text-xl font-bold">
          {gameId ? `Game: ${gameId}` : "Minesweeper Demo"}
      </div>
      {gameState ? (
        <MinesweeperBoard
          gameState={gameState}
          onCellClick={handleCellClick}
          onCellRightClick={handleCellRightClick}
        />
      ) : (
        <div class="text-white">Loading...</div>
      )}
    </div>
  );
}
