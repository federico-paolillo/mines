import { CellDto } from "../client/models/res";
import {
  LeftMouseButton,
  MiddleMouseButton,
  RightMouseButton,
} from "../utils/mouse";
import { MakeMoveEventHandler, useMoves } from "./use-moves";

export interface UseCellArgs {
  cell: CellDto;
  onMakeMove: MakeMoveEventHandler;
}

export function useCell({ cell, onMakeMove }: UseCellArgs) {
  const { open, chord, flag } = useMoves(cell, onMakeMove);

  function onClick(e: React.MouseEvent) {
    switch (e.button) {
      case LeftMouseButton:
        open();
        break;
      case MiddleMouseButton:
        chord();
        break;
      case RightMouseButton:
        flag();
        break;
    }
  }

  return {
    onClick,
  };
}
