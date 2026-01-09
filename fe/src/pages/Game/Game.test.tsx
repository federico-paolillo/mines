import { render, screen, waitFor } from "@testing-library/preact";
import { describe, expect, it, vi, beforeEach, afterEach, Mock } from "vitest";
import { Game } from "./index";
import { useRoute, useLocation } from "preact-iso";
import { useApiClient } from "../../clientContext";

// Mock dependencies
vi.mock("preact-iso", () => ({
  useRoute: vi.fn(),
  useLocation: vi.fn(),
}));

vi.mock("../../clientContext", () => ({
  useApiClient: vi.fn(),
}));

vi.mock("../../components/Minesweeper", () => ({
  MinesweeperBoard: () => <div data-testid="minesweeper-board">Board</div>,
}));

vi.mock("../../components/Spinner", () => ({
  Spinner: ({ isOpen }: { isOpen: boolean }) => (isOpen ? <div data-testid="spinner">Loading...</div> : null),
}));

vi.mock("../../components/Countdown", () => ({
  Countdown: ({ startTime, onExpired }: { startTime: number; onExpired: () => void }) => (
    <div data-testid="countdown">
      Countdown: {startTime}
      <button onClick={onExpired} data-testid="expire-button">Expire</button>
    </div>
  ),
}));

describe("Game Page", () => {
  const mockRoute = vi.fn();
  const mockFetchMatch = vi.fn();

  beforeEach(() => {
    (useLocation as Mock).mockReturnValue({ route: mockRoute });
    (useApiClient as Mock).mockReturnValue({
      fetchMatch: mockFetchMatch,
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should render loading state initially", () => {
    (useRoute as Mock).mockReturnValue({ query: "id=123" });
    mockFetchMatch.mockReturnValue(new Promise(() => {})); // Never resolves

    render(<Game />);
    expect(screen.getByTestId("spinner")).toBeInTheDocument();
  });

  it("should render game board and countdown when loaded", async () => {
    (useRoute as Mock).mockReturnValue({ query: "id=123" });
    const startTime = 1600000000;
    mockFetchMatch.mockResolvedValue({
      success: true,
      value: {
        id: "123",
        startTime: startTime,
        cells: [],
        width: 10,
        height: 10,
      },
    });

    render(<Game />);

    await waitFor(() => {
      expect(screen.queryByTestId("spinner")).not.toBeInTheDocument();
    });

    expect(screen.getByTestId("minesweeper-board")).toBeInTheDocument();
    expect(screen.getByTestId("countdown")).toHaveTextContent(`Countdown: ${startTime}`);
  });

  it("should navigate to game over when countdown expires", async () => {
    (useRoute as Mock).mockReturnValue({ query: "id=123" });
    const startTime = 1600000000;
    mockFetchMatch.mockResolvedValue({
      success: true,
      value: {
        id: "123",
        startTime: startTime,
        cells: [],
        width: 10,
        height: 10,
      },
    });

    render(<Game />);

    await waitFor(() => {
      expect(screen.queryByTestId("spinner")).not.toBeInTheDocument();
    });

    const expireButton = screen.getByTestId("expire-button");
    expireButton.click();

    expect(mockRoute).toHaveBeenCalledWith("/game-over");
  });

  it("should display error message if game load fails", async () => {
    (useRoute as Mock).mockReturnValue({ query: "id=123" });
    mockFetchMatch.mockResolvedValue({
      success: false,
      error: "Failed to fetch",
    });

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});

    render(<Game />);

    await waitFor(() => {
      expect(screen.queryByTestId("spinner")).not.toBeInTheDocument();
    });

    expect(screen.getByText("Game not found or failed to load.")).toBeInTheDocument();
    consoleSpy.mockRestore();
  });
});
