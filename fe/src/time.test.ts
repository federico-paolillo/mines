import { describe, expect, it } from "vitest";
import { toUnixTimestamp } from "./time";

describe("time", () => {
  it("should convert Date to UNIX timestamp in seconds", () => {
    const date = new Date("2020-01-01T00:00:00Z");
    const timestamp = toUnixTimestamp(date);
    expect(timestamp).toBe(1577836800);
  });

  it("should handle current time", () => {
    const now = new Date();
    const timestamp = toUnixTimestamp(now);
    expect(timestamp).toBe(Math.floor(now.getTime() / 1000));
  });
});
