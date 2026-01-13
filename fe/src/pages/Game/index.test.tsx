import { DefaultApiError } from "@microsoft/kiota-abstractions";
import { fireEvent, render, screen, waitFor } from "@testing-library/preact";
import { LocationProvider } from "preact-iso";
import { beforeEach, describe, expect, it, vi } from "vitest";
import type { MinesApiClient } from "../../api";
import { CellstateObject } from "../../client/models/board";
import { MovetypeObject } from "../../client/models/matchmaking";
import { ClientContext } from "../../clientContext";
import { Game } from "./index";

// Mock dependencies
const mockRoute = vi.fn();
vi.mock("preact-iso", async () => {
  const actual = await vi.importActual("preact-iso");
  return {
    ...actual,
    useLocation: () => ({ route: mockRoute }),
    useRoute: () => ({ query: { id: "test-game-id" } }),
  };
});

describe("Game Page", () => {
  let mockClient: MinesApiClient;

  beforeEach(() => {
    vi.clearAllMocks();
    mockClient = {
      fetchMatch: vi.fn(),
      startNewGame: vi.fn(),
      makeMove: vi.fn(),
    } as unknown as MinesApiClient;
  });

  const renderGame = () => {
    return render(
      <LocationProvider>
        <ClientContext.Provider value={mockClient}>
          <Game />
        </ClientContext.Provider>
      </LocationProvider>,
    );
  };

  const mockGameState = {
    id: "test-game-id",
    width: 2,
    height: 2,
    lives: 3,
    cells: [
      { x: 1, y: 1, state: CellstateObject.Closed },
      { x: 2, y: 1, state: CellstateObject.Closed },
      { x: 1, y: 2, state: CellstateObject.Closed },
      { x: 2, y: 2, state: CellstateObject.Closed },
    ],
    state: "playing",
  };

  it("fetches and displays game state on load", async () => {
    (mockClient.fetchMatch as any).mockResolvedValue({
      success: true,
      value: mockGameState,
    });

    renderGame();

    expect(mockClient.fetchMatch).toHaveBeenCalledWith("test-game-id");
    await waitFor(() => {
      expect(screen.getAllByRole("button")).toHaveLength(4);
    });
  });

  it("handles cell click and calls makeMove", async () => {
    (mockClient.fetchMatch as any).mockResolvedValue({
      success: true,
      value: mockGameState,
    });

    const updatedGameState = {
      ...mockGameState,
      cells: [
        { x: 1, y: 1, state: CellstateObject.Open, adjacentMines: 1 },
        ...mockGameState.cells.slice(1),
      ],
    };

    (mockClient.makeMove as any).mockResolvedValue({
      success: true,
      value: updatedGameState,
    });

    renderGame();

    await waitFor(() => screen.getAllByRole("button"));

    const cell = screen.getAllByRole("button")[0]; // x=1, y=1
    fireEvent.click(cell);

    expect(mockClient.makeMove).toHaveBeenCalledWith("test-game-id", {
      x: 1,
      y: 1,
      type: MovetypeObject.Open,
    });

    await waitFor(() => {
      expect(screen.getByText("1")).not.toBeNull();
    });
  });

  it("redirects to game over if lives are 0 after move", async () => {
    (mockClient.fetchMatch as any).mockResolvedValue({
      success: true,
      value: mockGameState,
    });

    const lostGameState = {
      ...mockGameState,
      lives: 0,
      state: "lost",
    };

    (mockClient.makeMove as any).mockResolvedValue({
      success: true,
      value: lostGameState,
    });

    renderGame();

    await waitFor(() => screen.getAllByRole("button"));
    fireEvent.click(screen.getAllByRole("button")[0]);

    await waitFor(() => {
      expect(mockRoute).toHaveBeenCalledWith("/game-over");
    });
  });

  it("redirects to game over on 422 error", async () => {
    (mockClient.fetchMatch as any).mockResolvedValue({
      success: true,
      value: mockGameState,
    });

    const error = new DefaultApiError("Match over");
    error.responseStatusCode = 422;

    (mockClient.makeMove as any).mockResolvedValue({
      success: false,
      error: { cause: error },
    });

    renderGame();

    await waitFor(() => screen.getAllByRole("button"));
    fireEvent.click(screen.getAllByRole("button")[0]);

    await waitFor(() => {
      expect(mockRoute).toHaveBeenCalledWith("/game-over");
    });
  });

  it("logs error on 409 and does not redirect", async () => {
    (mockClient.fetchMatch as any).mockResolvedValue({
      success: true,
      value: mockGameState,
    });

    const error = new DefaultApiError("Concurrent update");
    error.responseStatusCode = 409;
    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});

    (mockClient.makeMove as any).mockResolvedValue({
      success: false,
      error: { cause: error },
    });

    renderGame();

    await waitFor(() => screen.getAllByRole("button"));
    fireEvent.click(screen.getAllByRole("button")[0]);

    await waitFor(() => {
      expect(consoleSpy).toHaveBeenCalled();
    });
    expect(mockRoute).not.toHaveBeenCalled();
    consoleSpy.mockRestore();
  });
});
