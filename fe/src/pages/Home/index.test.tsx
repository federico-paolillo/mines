import { fireEvent, render, screen, waitFor } from "@testing-library/preact";
import { describe, expect, it, vi } from "vitest";
import { Home } from "./index";

// Mock dependencies
const mockStartNewGame = vi.fn();
const mockRoute = vi.fn();

vi.mock("../../clientContext", () => ({
  useApiClient: () => ({
    startNewGame: mockStartNewGame,
  }),
}));

vi.mock("preact-iso", () => ({
  useLocation: () => ({
    route: mockRoute,
  }),
}));

// Mock Spinner to easily check if it's rendered
vi.mock("../../components/Spinner", () => ({
  Spinner: ({ isOpen }: { isOpen: boolean }) =>
    isOpen ? <div data-testid="spinner">Loading...</div> : null,
}));

describe("Home", () => {
  it("shows spinner while loading and navigates on success", async () => {
    let resolvePromise: (value: unknown) => void = () => {};

    const promise = new Promise((resolve) => {
      resolvePromise = resolve;
    });

    mockStartNewGame.mockReturnValue(promise);

    render(<Home />);

    fireEvent.click(screen.getByText("New Game"));

    expect(screen.getByTestId("spinner")).not.toBeNull();

    resolvePromise({
      success: true,
      value: { id: "new-game-id" },
    });

    await waitFor(() =>
      expect(mockRoute).toHaveBeenCalledWith("/game?id=new-game-id"),
    );

    expect(mockStartNewGame).toHaveBeenCalled();
  });
});
