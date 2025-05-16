import { CellDto } from "../client/models/res";
import { Cell } from "./Cell";

export interface BoardProps {
  width: number;
  height: number;
  cells: CellDto[];
}

export function Board({ width, height, cells }: BoardProps) {
  const grid = asGrid(width, height, cells);

  return (
    <div className="flex flex-col gap-0">
      {grid.map((row, rowIndex) => (
        <div className="flex flex-row gap-0" key={`row-${rowIndex}`}>
          {row.map((cell, colIndex) => (
            <Cell cell={cell} key={`${colIndex}`} onMakeMove={() => {}}></Cell>
          ))}
        </div>
      ))}
    </div>
  );
}

function asGrid(width: number, height: number, cells: CellDto[]): CellDto[][] {
  const grid: CellDto[][] = [];

  for (let row = 0; row < height; row++) {
    grid[row] = [];
    for (let col = 0; col < width; col++) {
      const index1d = col + row * width;
      grid[row][col] = cells[index1d];
    }
  }

  return grid;
}
