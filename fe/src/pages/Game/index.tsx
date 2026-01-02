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
    if (gameId) {
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
        id: gameId,
        lives: 3,
        state: GamestateObject.Playing,
        width,
        height,
        cells,
      });
    } else {
        // Just for demo if no ID provided
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
         id: "demo",
         lives: 3,
         state: GamestateObject.Playing,
         width,
         height,
         cells,
       });
    }
  }, [gameId]);

  const handleCellClick = (x: number, y: number) => {
    console.log(`Clicked cell at ${x}, ${y}`);
    // Here we would call the API to open the cell
    // For visual feedback in this demo, let's toggle state locally
    setGameState((prev) => {
      if (!prev || !prev.cells) return prev;
      const newCells = prev.cells.map((c) => {
        if (c.x === x && c.y === y) {
          return { ...c, state: CellstateObject.Open };
        }
        return c;
      });
      return { ...prev, cells: newCells };
    });
  };

  const handleCellRightClick = (x: number, y: number) => {
    console.log(`Right clicked cell at ${x}, ${y}`);
    // Here we would call the API to flag the cell
    setGameState((prev) => {
        if (!prev || !prev.cells) return prev;
        const newCells = prev.cells.map((c) => {
          if (c.x === x && c.y === y) {
             const newState = c.state === CellstateObject.Closed ? CellstateObject.Flagged : (c.state === CellstateObject.Flagged ? CellstateObject.Closed : c.state);
            return { ...c, state: newState };
          }
          return c;
        });
        return { ...prev, cells: newCells };
      });
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
