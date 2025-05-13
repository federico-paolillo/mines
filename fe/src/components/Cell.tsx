import { CellDto } from "../client/models/res";
import { Flag } from "./icons/Flag";
import { Square } from "./icons/Square";
import { Warn } from "./icons/Warn";

export interface CellProps {
  cell: CellDto;
}

export function Cell({ cell }: CellProps) {
  switch (cell.state) {
    case "open":
      return <Cell cell={cell} />;
    case "closed":
      return <Square />;
    case "flagged":
      return <Flag />;
    case "unfathomable":
      return <Warn />;
  }
}
