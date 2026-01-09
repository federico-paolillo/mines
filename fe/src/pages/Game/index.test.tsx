import { render, screen, waitFor } from "@testing-library/preact";
import { describe, expect, it, vi } from "vitest";
import { Game } from "./index";

// Mock dependencies
const mockFetchMatch = vi.fn();
const mockRouteQuery = { id: "test-game-id" };

vi.mock("../../clientContext", () => ({
  useApiClient: () => ({
    fetchMatch: mockFetchMatch,
  }),
}));

vi.mock("preact-iso", () => ({
  useRoute: () => ({
    query: mockRouteQuery,
  }),
  useLocation: () => ({
    route: vi.fn(),
  }),
}));

vi.mock("../../components/Spinner", () => ({
  Spinner: ({ isOpen }: { isOpen: boolean }) =>
    isOpen ? <div data-testid="spinner">Loading...</div> : null,
}));

vi.mock("../../components/Minesweeper", () => ({
  MinesweeperBoard: () => <div data-testid="board">Board</div>,
}));

describe("Game", () => {
  it("fetches game data and shows spinner", async () => {
    // Return a promise that doesn't resolve immediately to check spinner
    let resolvePromise: (value: unknown) => void = () => {};

    const promise = new Promise((resolve) => {
      resolvePromise = resolve;
    });

    mockFetchMatch.mockReturnValue(promise);

    render(<Game />);

    // Check spinner is visible
    expect(screen.getByTestId("spinner")).not.toBeNull();

    // Resolve promise
    resolvePromise({
      success: true,
      value: {
        id: "test-game-id",
        cells: [],
        width: 9,
        height: 9,
        state: "Playing",
      },
    });

    // Wait for spinner to disappear and board to appear
    await waitFor(() => expect(screen.queryByTestId("spinner")).toBeNull());
    expect(screen.getByTestId("board")).not.toBeNull();
    expect(mockFetchMatch).toHaveBeenCalledWith("test-game-id");
  });
});
