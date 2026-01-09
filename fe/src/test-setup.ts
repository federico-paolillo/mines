import { cleanup } from "@testing-library/preact";
import { afterEach, beforeEach, vi } from "vitest";

beforeEach(() => {
  vi.clearAllMocks();
});

afterEach(() => {
  cleanup();
});
