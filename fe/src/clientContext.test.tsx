import { render } from "@testing-library/preact";
import { describe, expect, test, vi } from "vitest";
import type { MinesApiClient } from "./api";
import { ClientProvider, useApiClient, withErrorLogging } from "./clientContext";
import { success, failure } from "./result";

test("useApiClient returns client within ClientProvider", () => {
  let client: MinesApiClient | undefined;

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

describe("withErrorLogging", () => {
  const mockClient: MinesApiClient = {
    fetchMatch: vi.fn(),
    startNewGame: vi.fn(),
    makeMove: vi.fn(),
  };

  test("passes through successful results without logging", async () => {
    const value = { id: "abc", lives: 3 };
    vi.mocked(mockClient.fetchMatch).mockResolvedValue(success(value));

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = withErrorLogging(mockClient);

    const result = await wrapped.fetchMatch("abc");

    expect(result).toEqual(success(value));
    expect(consoleSpy).not.toHaveBeenCalled();

    consoleSpy.mockRestore();
  });

  test("logs to console.error on failure and returns the failure", async () => {
    const cause = new Error("network down");
    vi.mocked(mockClient.makeMove).mockResolvedValue(
      failure({ message: "Move failed", cause }),
    );

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = withErrorLogging(mockClient);

    const result = await wrapped.makeMove("abc", { x: 0, y: 0 } as any);

    expect(result.success).toBe(false);
    expect(consoleSpy).toHaveBeenCalledWith("Move failed", cause);

    consoleSpy.mockRestore();
  });

  test("logs to console.error on failure without cause", async () => {
    vi.mocked(mockClient.startNewGame).mockResolvedValue(
      failure({ message: "Server error" }),
    );

    const consoleSpy = vi.spyOn(console, "error").mockImplementation(() => {});
    const wrapped = withErrorLogging(mockClient);

    const result = await wrapped.startNewGame({ difficulty: "beginner" } as any);

    expect(result.success).toBe(false);
    expect(consoleSpy).toHaveBeenCalledWith("Server error", undefined);

    consoleSpy.mockRestore();
  });
});
