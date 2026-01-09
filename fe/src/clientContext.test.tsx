import { render } from "@testing-library/preact";
import { expect, test } from "vitest";
import { ClientProvider, useApiClient } from "./clientContext";

test("useApiClient returns client within ClientProvider", () => {
  let client;
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

  // Prevent console.error from cluttering the output (React/Preact logs errors for caught exceptions in boundaries)
  const originalConsoleError = console.error;
  console.error = () => {};

  try {
    expect(() => render(<TestComponent />)).toThrow(
      "useApiClient must be used within a ClientProvider",
    );
  } finally {
    console.error = originalConsoleError;
  }
});
