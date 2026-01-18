import { renderHook, waitFor } from "@testing-library/preact";
import { describe, expect, it, vi } from "vitest";
import { useGameState } from "./useGameState";
import { ClientContext } from "../clientContext";
import { MovetypeObject } from "../client/models/matchmaking";
import { success, failure } from "../result";
import { DefaultApiError } from "@microsoft/kiota-abstractions";

// Mock useLocation
const mockRoute = vi.fn();
vi.mock("preact-iso", () => ({
  useLocation: () => ({ route: mockRoute }),
}));

describe("useGameState", () => {
  const mockClient = {
    fetchMatch: vi.fn(),
    makeMove: vi.fn(),
  };

  const wrapper = ({ children }: { children: any }) => (
    <ClientContext.Provider value={mockClient as any}>
      {children}
    </ClientContext.Provider>
  );

  it("should fetch game state on mount", async () => {
    const gameId = "game-123";
    const mockGameState = {
      lives: 3,
      startTime: 1234567890,
      board: [],
    };

    mockClient.fetchMatch.mockResolvedValue(success(mockGameState));

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    expect(result.current.loading).toBe(true);

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(mockClient.fetchMatch).toHaveBeenCalledWith(gameId);
    expect(result.current.gameState).toEqual(mockGameState);
  });

  it("should handle error when fetching game fails", async () => {
    const gameId = "game-123";
    mockClient.fetchMatch.mockResolvedValue(
      failure({ message: "Error fetching game" })
    );

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.gameState).toBeNull();
    // Error logging is mocked/console.error, not asserting on console.error here but verifying state
  });

  it("should redirect to game-over if lives are 0 on fetch", async () => {
    const gameId = "game-123";
    const mockGameState = {
      lives: 0,
      startTime: 1234567890,
      board: [],
    };

    mockClient.fetchMatch.mockResolvedValue(success(mockGameState));

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(mockRoute).toHaveBeenCalledWith("/game-over");
    });
  });

  it("should handle cell click (make move)", async () => {
    const gameId = "game-123";
    const initialGameState = {
      lives: 3,
      startTime: 1234567890,
      board: [],
    };
    const updatedGameState = {
      lives: 3,
      startTime: 1234567890,
      board: ["open"],
    };

    mockClient.fetchMatch.mockResolvedValue(success(initialGameState));
    mockClient.makeMove.mockResolvedValue(success(updatedGameState));

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    await result.current.onCellClick(1, 1);

    expect(mockClient.makeMove).toHaveBeenCalledWith(gameId, {
      x: 1,
      y: 1,
      type: MovetypeObject.Open,
    });

    await waitFor(() => {
        expect(result.current.gameState).toEqual(updatedGameState);
    });
  });

  it("should handle cell right click (flag)", async () => {
    const gameId = "game-123";
    const initialGameState = {
      lives: 3,
      startTime: 1234567890,
      board: [],
    };
    const updatedGameState = {
      lives: 3,
      startTime: 1234567890,
      board: ["flagged"],
    };

    mockClient.fetchMatch.mockResolvedValue(success(initialGameState));
    mockClient.makeMove.mockResolvedValue(success(updatedGameState));

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    await result.current.onCellRightClick(2, 2);

    expect(mockClient.makeMove).toHaveBeenCalledWith(gameId, {
      x: 2,
      y: 2,
      type: MovetypeObject.Flag,
    });

    await waitFor(() => {
        expect(result.current.gameState).toEqual(updatedGameState);
    });
  });

  it("should redirect to game-over if lives become 0 after move", async () => {
    const gameId = "game-123";
    const initialGameState = {
      lives: 1,
      startTime: 1234567890,
      board: [],
    };
    const updatedGameState = {
      lives: 0,
      startTime: 1234567890,
      board: [],
    };

    mockClient.fetchMatch.mockResolvedValue(success(initialGameState));
    mockClient.makeMove.mockResolvedValue(success(updatedGameState));

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    await result.current.onCellClick(1, 1);

    expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });

  it("should redirect to game-over on 422 error during move", async () => {
    const gameId = "game-123";
    const initialGameState = {
        lives: 1,
        startTime: 1234567890,
        board: [],
    };

    mockClient.fetchMatch.mockResolvedValue(success(initialGameState));

    const apiError = new Error("422 Error");
    (apiError as any).responseStatusCode = 422;
    const defaultApiError = new DefaultApiError("422 Error");
    defaultApiError.responseStatusCode = 422;

    mockClient.makeMove.mockResolvedValue(
        failure({ message: "Move failed", cause: defaultApiError })
    );

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
        expect(result.current.loading).toBe(false);
    });

    await result.current.onCellClick(1, 1);

    expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });

  it("should handle expired timer", () => {
      const { result } = renderHook(() => useGameState("game-123"), { wrapper });
      result.current.onExpired();
      expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });
});
