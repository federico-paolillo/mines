import { fireEvent, render, screen } from "@testing-library/preact";
import { describe, expect, it, vi } from "vitest";
import { CellstateObject } from "../../client/models/board";
import type { CellDto, MatchstateDto } from "../../client/models/res";
import { MinesweeperBoard } from "./Board";

describe("MinesweeperBoard Component", () => {
  const defaultGameState: MatchstateDto = {
    width: 3,
    height: 3,
    cells: [
      { x: 1, y: 1, state: CellstateObject.Closed },
      { x: 2, y: 1, state: CellstateObject.Closed },
      { x: 3, y: 1, state: CellstateObject.Closed },
      { x: 1, y: 2, state: CellstateObject.Closed },
      { x: 2, y: 2, state: CellstateObject.Closed },
      { x: 3, y: 2, state: CellstateObject.Closed },
      { x: 1, y: 3, state: CellstateObject.Closed },
      { x: 2, y: 3, state: CellstateObject.Closed },
      { x: 3, y: 3, state: CellstateObject.Closed },
    ] as CellDto[],
  };

  it("renders loading state when data is incomplete", () => {
    const onOpenCell = vi.fn();
    const onFlagCell = vi.fn();

    render(
      <MinesweeperBoard
        gameState={{}}
        onOpenCell={onOpenCell}
        onFlagCell={onFlagCell}
      />,
    );

    expect(screen.getByText("Loading board...")).toBeDefined();
  });

  it("renders the correct grid of cells", () => {
    const onOpenCell = vi.fn();
    const onFlagCell = vi.fn();

    render(
      <MinesweeperBoard
        gameState={defaultGameState}
        onOpenCell={onOpenCell}
        onFlagCell={onFlagCell}
      />,
    );

    const cells = screen.getAllByRole("button");
    expect(cells.length).toBe(9); // 3x3 grid
  });

  it("handles cell clicks and propagates to onOpenCell", () => {
    const onOpenCell = vi.fn();
    const onFlagCell = vi.fn();

    render(
      <MinesweeperBoard
        gameState={defaultGameState}
        onOpenCell={onOpenCell}
        onFlagCell={onFlagCell}
      />,
    );

    const cells = screen.getAllByRole("button");
    // Click the first cell (1, 1)
    fireEvent.click(cells[0]);
    expect(onOpenCell).toHaveBeenCalledWith(1, 1);

    // Click the middle cell (2, 2) which should be index 4
    fireEvent.click(cells[4]);
    expect(onOpenCell).toHaveBeenCalledWith(2, 2);
  });

  it("handles cell right-clicks and propagates to onFlagCell", () => {
    const onOpenCell = vi.fn();
    const onFlagCell = vi.fn();

    render(
      <MinesweeperBoard
        gameState={defaultGameState}
        onOpenCell={onOpenCell}
        onFlagCell={onFlagCell}
      />,
    );

    const cells = screen.getAllByRole("button");
    // Right-click the last cell (3, 3) which should be index 8
    fireEvent.contextMenu(cells[8]);
    expect(onFlagCell).toHaveBeenCalledWith(3, 3, expect.any(MouseEvent));
  });
});
