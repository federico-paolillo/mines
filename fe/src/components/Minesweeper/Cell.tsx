import { type Cellstate, CellstateObject } from "../../client/models/board";
import type { CellDto } from "../../client/models/res";
import { maybeToString } from "../../strings";

interface CellProps {
  cell: CellDto;
  onClick: (x: number, y: number) => void;
  onContextMenu: (x: number, y: number, e: MouseEvent) => void;
}

const baseClasses =
  "w-6 h-6 flex items-center justify-center text-xs font-bold select-none cursor-default";

function getCellStyle(state?: Cellstate | null): string {
  switch (state) {
    case CellstateObject.Closed:
      return "bg-[#c0c0c0] border-t-2 border-l-2 border-white border-b-2 border-r-2 border-[#808080]";
    case CellstateObject.Open:
      return "bg-[#c0c0c0] border border-[#808080]";
    case CellstateObject.Flagged:
      return "bg-[#c0c0c0] border-t-2 border-l-2 border-white border-b-2 border-r-2 border-[#808080]";
    case CellstateObject.Unfathomable:
      return "bg-red-600 border border-[#808080]";
    default:
      return "bg-[#c0c0c0]";
  }
}

function getCellContent(cell: CellDto): string | null {
  switch (cell.state) {
    case CellstateObject.Flagged:
      return "🚩";
    case CellstateObject.Unfathomable:
      return "💣";
    case CellstateObject.Open:
      return maybeToString(cell.adjacentMines) || null;
    default:
      return null;
  }
}

export function Cell({ cell, onClick, onContextMenu }: CellProps) {
  const { state, x, y } = cell;

  const handleClick = () => {
    if (x && y) {
      onClick(x, y);
    }
  };

  const handleContextMenu = (e: MouseEvent) => {
    e.preventDefault();
    if (x && y) {
      onContextMenu(x, y, e);
    }
  };

  const specificClasses = getCellStyle(state);
  const content = getCellContent(cell);

  return (
    <div
      class={`${baseClasses} ${specificClasses}`}
      onClick={handleClick}
      onContextMenu={handleContextMenu}
      role="button"
      aria-label={`Cell at ${x}, ${y}, ${state}`}
    >
      {content}
    </div>
  );
}
