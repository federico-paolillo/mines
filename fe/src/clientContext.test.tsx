import { render } from "@testing-library/preact";
import { expect, test } from "vitest";
import type { MinesApiClient } from "./api";
import { ClientProvider, useApiClient } from "./clientContext";

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
