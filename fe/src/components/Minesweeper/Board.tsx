import { MatchstateDto, CellDto } from "../../client/models/res";
import { Cell } from "./Cell";

interface MinesweeperBoardProps {
  gameState: MatchstateDto;
  onCellClick: (x: number, y: number) => void;
  onCellRightClick: (x: number, y: number) => void;
}

export function MinesweeperBoard({
  gameState,
  onCellClick,
  onCellRightClick,
}: MinesweeperBoardProps) {
  const { cells, width, height } = gameState;

  if (!cells || !width || !height) {
    return <div>Loading board...</div>;
  }

  // Create a grid structure
  const grid: CellDto[][] = [];
  for (let i = 0; i < height; i++) {
    const row: CellDto[] = [];
    for (let j = 0; j < width; j++) {
      // Find cell at (j+1, i+1) because 1-based indexing in backend but let's check.
      // The CellDto has x and y. Let's rely on finding the cell with matching x,y.
      // However, iterating the cells array is O(N). Doing it for every cell is O(N^2).
      // A map would be better.
      row.push({ state: "closed", x: j + 1, y: i + 1 } as unknown as CellDto); // Placeholder
    }
    grid.push(row);
  }

  // Populate grid from flat cells array
  cells.forEach((cell) => {
    if (
      cell.x !== null &&
      cell.y !== null &&
      cell.x !== undefined &&
      cell.y !== undefined
    ) {
        // Backend seems to use 1-based indexing based on loop in export.go: location := dimensions.Location{X: x + 1, Y: y + 1}
        // So we adjust to 0-based for array access.
        const rowIdx = cell.y - 1;
        const colIdx = cell.x - 1;

        if (rowIdx >= 0 && rowIdx < height && colIdx >= 0 && colIdx < width) {
            grid[rowIdx][colIdx] = cell;
        }
    }
  });

  return (
    <div
      class="inline-block border-l-4 border-t-4 border-white border-r-4 border-b-4 border-[#808080] bg-[#c0c0c0] p-1"
      style={{
        display: "grid",
        gridTemplateColumns: `repeat(${width}, 1.5rem)`, // 1.5rem = 24px (w-6)
        gap: "0",
      }}
    >
      {grid.map((row, rowIndex) =>
        row.map((cell, colIndex) => (
          <Cell
            key={`${cell.x}-${cell.y}`}
            cell={cell}
            onClick={onCellClick}
            onContextMenu={(x, y, e) => onCellRightClick(x, y)}
          />
        ))
      )}
    </div>
  );
}
