import type { CellDto, MatchstateDto } from "../../client/models/res";
import { Cell } from "./Cell";

interface MinesweeperBoardProps {
  gameState: MatchstateDto;
  onOpenCell: (x: number, y: number) => void;
  onFlagCell: (x: number, y: number) => void;
}

function makeGrid(
  cells: CellDto[],
  width: number,
  height: number,
): CellDto[][] {
  const grid: CellDto[][] = [];

  for (let i = 0; i < height; i++) {
    const row: CellDto[] = [];

    for (let j = 0; j < width; j++) {
      row.push({ state: "closed", x: j + 1, y: i + 1 } as unknown as CellDto);
    }

    grid.push(row);
  }

  for (const cell of cells) {
    if (cell.x && cell.y) {
      const rowIdx = cell.y - 1;
      const colIdx = cell.x - 1;

      if (rowIdx >= 0 && rowIdx < height && colIdx >= 0 && colIdx < width) {
        grid[rowIdx][colIdx] = cell;
      } else {
        throw new Error(`Cell out of bounds: ${cell.x}, ${cell.y}`);
      }
    }
  }

  return grid;
}

export function MinesweeperBoard({
  gameState,
  onOpenCell,
  onFlagCell,
}: MinesweeperBoardProps) {
  const { cells, width, height } = gameState;

  if (!cells || !width || !height) {
    return <div>Loading board...</div>;
  }

  const grid = makeGrid(cells, width, height);

  return (
    <div
      class="inline-block border-l-4 border-t-4 border-r-4 border-b-4 border-[#808080] bg-[#c0c0c0] p-1"
      style={{
        display: "grid",
        gridTemplateColumns: `repeat(${width}, 1.5rem)`, // 1.5rem = 24px (w-6)
        gap: "0",
      }}
    >
      {grid.map((row) =>
        row.map((cell) => (
          <Cell
            key={`${cell.x}-${cell.y}`}
            cell={cell}
            onClick={onOpenCell}
            onContextMenu={onFlagCell}
          />
        )),
      )}
    </div>
  );
}
