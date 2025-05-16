import { useCallback } from "react";
import { Cellstate, CellstateObject } from "../client/models/board";
import { Movetype, MovetypeObject } from "../client/models/matchmaking";
import { CellDto } from "../client/models/res";
import { RequireAll } from "../utils/require-all";

const movesForState = new Map<Movetype, Cellstate[]>([
  [MovetypeObject.Flag, [CellstateObject.Closed]],
  [MovetypeObject.Open, [CellstateObject.Closed, CellstateObject.Flagged]],
  [MovetypeObject.Chord, [CellstateObject.Closed, CellstateObject.Flagged]],
]);

function canMakeMove(movetype: Movetype, cellstate: Cellstate) {
  const possibleStates = movesForState.get(movetype) ?? [];
  return possibleStates.includes(cellstate);
}

export type MakeMoveEventHandler = (
  moveType: Movetype,
  x: number,
  y: number
) => void;

export function useMoves(cell: CellDto, onMakeMove: MakeMoveEventHandler) {
  const { x, y, state } = cell;

  const open = useCallback(() => {
    if (canMakeMove(MovetypeObject.Open, state!)) {
      return onMakeMove(MovetypeObject.Open, x!, y!);
    }
  }, [onMakeMove, state, x, y]);

  const chord = useCallback(() => {
    if (canMakeMove(MovetypeObject.Chord, state!)) {
      return onMakeMove(MovetypeObject.Chord, x!, y!);
    }
  }, [onMakeMove, state, x, y]);

  const flag = useCallback(() => {
    if (canMakeMove(MovetypeObject.Flag, state!)) {
      return onMakeMove(MovetypeObject.Flag, x!, y!);
    }
  }, [onMakeMove, state, x, y]);

  return {
    open,
    chord,
    flag,
  };
}
