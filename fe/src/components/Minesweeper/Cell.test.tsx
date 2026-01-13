import { cleanup, fireEvent, render, screen } from "@testing-library/preact";
import { afterEach, describe, expect, it, vi } from "vitest";
import { CellstateObject } from "../../client/models/board";
import type { CellDto } from "../../client/models/res";
import { Cell } from "./Cell";

describe("Cell Component", () => {
  afterEach(() => {
    cleanup();
  });

  const defaultCell: CellDto = {
    x: 1,
    y: 1,
    state: CellstateObject.Closed,
  };

  it("renders closed cell correctly", () => {
    const onClick = vi.fn();
    const onContextMenu = vi.fn();

    render(
      <Cell
        cell={{ ...defaultCell, state: CellstateObject.Closed }}
        onClick={onClick}
        onContextMenu={onContextMenu}
      />,
    );

    const cellElement = screen.getByRole("button");
    expect(cellElement).toBeDefined();
    expect(cellElement.getAttribute("aria-label")).toContain("closed");
    expect(screen.getByText("?")).toBeDefined();
  });

  it("renders flagged cell correctly", () => {
    render(
      <Cell
        cell={{ ...defaultCell, state: CellstateObject.Flagged }}
        onClick={vi.fn()}
        onContextMenu={vi.fn()}
      />,
    );

    expect(screen.getByText("🚩")).toBeDefined();
  });

  it("renders unfathomable (bomb) cell correctly", () => {
    render(
      <Cell
        cell={{ ...defaultCell, state: CellstateObject.Unfathomable }}
        onClick={vi.fn()}
        onContextMenu={vi.fn()}
      />,
    );

    expect(screen.getByText("💣")).toBeDefined();
  });

  it("calls onClick handler when clicked", () => {
    const onClick = vi.fn();
    const onContextMenu = vi.fn();

    render(
      <Cell
        cell={defaultCell}
        onClick={onClick}
        onContextMenu={onContextMenu}
      />,
    );

    const cellElement = screen.getByRole("button");
    fireEvent.click(cellElement);

    expect(onClick).toHaveBeenCalledTimes(1);
    expect(onClick).toHaveBeenCalledWith(1, 1);
  });

  it("calls onContextMenu handler when right-clicked", () => {
    const onClick = vi.fn();
    const onContextMenu = vi.fn();

    render(
      <Cell
        cell={defaultCell}
        onClick={onClick}
        onContextMenu={onContextMenu}
      />,
    );

    const cellElement = screen.getByRole("button");
    fireEvent.contextMenu(cellElement);

    expect(onContextMenu).toHaveBeenCalledTimes(1);
    expect(onContextMenu).toHaveBeenCalledWith(1, 1, expect.any(Object));
  });

  it("does not call handlers if x or y are missing", () => {
    const onClick = vi.fn();
    const onContextMenu = vi.fn();
    const incompleteCell: CellDto = { state: CellstateObject.Closed };

    render(
      <Cell
        cell={incompleteCell}
        onClick={onClick}
        onContextMenu={onContextMenu}
      />,
    );

    const cellElement = screen.getByRole("button");
    fireEvent.click(cellElement);
    fireEvent.contextMenu(cellElement);

    expect(onClick).not.toHaveBeenCalled();
    expect(onContextMenu).not.toHaveBeenCalled();
  });
});
