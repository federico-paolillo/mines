import { CellDto } from "../../client/models/res";
import { Cellstate, CellstateObject } from "../../client/models/board";

interface CellProps {
  cell: CellDto;
  onClick: (x: number, y: number) => void;
  onContextMenu: (x: number, y: number, e: MouseEvent) => void;
}

export function Cell({ cell, onClick, onContextMenu }: CellProps) {
  const { state, x, y } = cell;

  const handleClick = () => {
    if (x !== undefined && y !== undefined && x !== null && y !== null) {
      onClick(x, y);
    }
  };

  const handleContextMenu = (e: MouseEvent) => {
    e.preventDefault();
    if (x !== undefined && y !== undefined && x !== null && y !== null) {
      onContextMenu(x, y, e);
    }
  };

  const baseClasses =
    "w-6 h-6 flex items-center justify-center text-xs font-bold select-none cursor-default";

  let specificClasses = "";
  let content = null;

  switch (state) {
    case CellstateObject.Closed:
      specificClasses =
        "bg-[#c0c0c0] border-t-2 border-l-2 border-white border-b-2 border-r-2 border-[#808080]";
      break;
    case CellstateObject.Open:
      // Depressed look: border-t-gray, border-l-gray, border-b-white? No, standard win95 is flat or simple thin border.
      // Actually often just flat gray with no bevel, or a 1px darker border.
      // Let's go with a simple flat look or thin border.
      specificClasses = "bg-[#c0c0c0] border border-[#808080]";

      // Note: The API currently does not return the number of adjacent mines.
      // If it did, we would render it here with specific colors.
      // 1: Blue, 2: Green, 3: Red, 4: Dark Blue, 5: Maroon, 6: Turquoise, 7: Black, 8: Gray
      break;
    case CellstateObject.Flagged:
      specificClasses =
        "bg-[#c0c0c0] border-t-2 border-l-2 border-white border-b-2 border-r-2 border-[#808080]";
      content = "🚩"; // Or use an SVG or FontAwesome icon if available
      break;
    case CellstateObject.Unfathomable:
      // This usually means a mine was revealed (game over)
      specificClasses = "bg-red-600 border border-[#808080]";
      content = "💣";
      break;
    default:
      specificClasses = "bg-[#c0c0c0]";
  }

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
