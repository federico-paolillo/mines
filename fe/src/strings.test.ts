import { describe, expect, it } from "vitest";
import { padWithTwoZeros } from "./strings";

describe("strings", () => {
  it("should pad single digit number string", () => {
    expect(padWithTwoZeros("5")).toBe("05");
  });

  it("should not pad double digit number string", () => {
    expect(padWithTwoZeros("12")).toBe("12");
  });

  it("should not pad longer strings", () => {
    expect(padWithTwoZeros("123")).toBe("123");
  });

  it("should pad empty string", () => {
    expect(padWithTwoZeros("")).toBe("00");
  });
});
