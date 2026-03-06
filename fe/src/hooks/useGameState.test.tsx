import { renderHook, waitFor } from "@testing-library/preact";
import { describe, expect, it, vi } from "vitest";
import { useGameState } from "./useGameState";
import { ClientContext } from "../clientContext";
import { MovetypeObject } from "../client/models/matchmaking";
import { type ApiError, success } from "../result";

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
    const error: ApiError = { kind: "unknown", message: "Error fetching game" };
    mockClient.fetchMatch.mockResolvedValue({ success: false, error });

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.gameState).toBeNull();
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

    await result.current.onOpenCell(1, 1);

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

    await result.current.onFlagCell(2, 2);

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

    await result.current.onOpenCell(1, 1);

    expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });

  it("should redirect to game-over on match_over error during move", async () => {
    const gameId = "game-123";
    const initialGameState = {
        lives: 1,
        startTime: 1234567890,
        board: [],
    };

    mockClient.fetchMatch.mockResolvedValue(success(initialGameState));

    const error: ApiError = { kind: "match_over", message: "Move failed" };
    mockClient.makeMove.mockResolvedValue({ success: false, error });

    const { result } = renderHook(() => useGameState(gameId), { wrapper });

    await waitFor(() => {
        expect(result.current.loading).toBe(false);
    });

    await result.current.onOpenCell(1, 1);

    expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });

  it("should handle expired timer", () => {
      const { result } = renderHook(() => useGameState("game-123"), { wrapper });
      result.current.onExpired();
      expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });
});
