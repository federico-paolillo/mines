import { DefaultApiError } from "@microsoft/kiota-abstractions";
import { render } from "@testing-library/preact";
import { describe, expect, test, vi } from "vitest";
import type { MinesApiClient } from "./api";
import {
  ClientProvider,
  type GameClient,
  useApiClient,
  wrapClient,
} from "./clientContext";
import { failure, success } from "./result";

test("useApiClient returns client within ClientProvider", () => {
  let client: GameClient | undefined;

  const TestComponent = () => {
    client = useApiClient();
    return null;
  };

  render(
    <ClientProvider>
      <TestComponent />
    </ClientProvider>,
  );

  expect(client).toBeDefined();
});

test("useApiClient throws error outside ClientProvider", () => {
  const TestComponent = () => {
    useApiClient();
    return null;
  };

  expect(() => render(<TestComponent />)).toThrow(
    "useApiClient must be used within a ClientProvider",
  );
});

describe("wrapClient", () => {
  const mockClient: MinesApiClient = {
    fetchMatch: vi.fn(),
    startNewGame: vi.fn(),
    makeMove: vi.fn(),
  };

  test("passes through successful results without logging", async () => {
    const value = { id: "abc", lives: 3 };
    vi.mocked(mockClient.fetchMatch).mockResolvedValue(success(value));

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = wrapClient(mockClient);

    const result = await wrapped.fetchMatch("abc");

    expect(result).toEqual(success(value));
    expect(consoleSpy).not.toHaveBeenCalled();

    consoleSpy.mockRestore();
  });

  test("logs to console.error on failure and returns ApiError with unknown kind", async () => {
    const cause = new Error("network down");
    vi.mocked(mockClient.makeMove).mockResolvedValue(
      failure({ message: "Move failed", cause }),
    );

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = wrapClient(mockClient);

    const result = await wrapped.makeMove("abc", { x: 0, y: 0 } as any);

    expect(result.success).toBe(false);
    if (!result.success) {
      expect(result.error).toEqual({ kind: "unknown", message: "Move failed" });
    }
    expect(consoleSpy).toHaveBeenCalledWith("Move failed", cause);

    consoleSpy.mockRestore();
  });

  test("translates 422 DefaultApiError to match_over kind", async () => {
    const apiError = new DefaultApiError("Match over");
    apiError.responseStatusCode = 422;

    vi.mocked(mockClient.makeMove).mockResolvedValue(
      failure({ message: "Move failed", cause: apiError }),
    );

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = wrapClient(mockClient);

    const result = await wrapped.makeMove("abc", { x: 0, y: 0 } as any);

    expect(result.success).toBe(false);
    if (!result.success) {
      expect(result.error).toEqual({
        kind: "match_over",
        message: "Move failed",
      });
    }
    expect(consoleSpy).toHaveBeenCalledWith("Move failed", apiError);

    consoleSpy.mockRestore();
  });

  test("translates non-422 DefaultApiError to unknown kind", async () => {
    const apiError = new DefaultApiError("Conflict");
    apiError.responseStatusCode = 409;

    vi.mocked(mockClient.fetchMatch).mockResolvedValue(
      failure({ message: "Fetch failed", cause: apiError }),
    );

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = wrapClient(mockClient);

    const result = await wrapped.fetchMatch("abc");

    expect(result.success).toBe(false);
    if (!result.success) {
      expect(result.error).toEqual({
        kind: "unknown",
        message: "Fetch failed",
      });
    }

    consoleSpy.mockRestore();
  });
});
