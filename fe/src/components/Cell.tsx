import { CellDto } from "../client/models/res";
import { useCell } from "../hooks/use-cell";
import { MakeMoveEventHandler } from "../hooks/use-moves";
import { Flag } from "./icons/Flag";
import { Number } from "./icons/Number";
import { Square } from "./icons/Square";
import { Warn } from "./icons/Warn";

export interface CellProps {
  cell: CellDto;
  onMakeMove: MakeMoveEventHandler;
}

export function Cell({ cell, onMakeMove }: CellProps) {
  const { onClick } = useCell({
    cell,
    onMakeMove,
  });

  return <div onClick={onClick}>{renderCell(cell)}</div>;
}

function renderCell(cell: CellDto) {
  switch (cell.state) {
    case "open":
      return <Number val={1} />;
    case "closed":
      return <Square />;
    case "flagged":
      return <Flag />;
    default:
      return <Warn />;
  }
}
