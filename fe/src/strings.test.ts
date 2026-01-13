import { describe, expect, it } from "vitest";
import { maybeToString, padWithTwoZeros } from "./strings";

describe("strings", () => {
  describe("padWithTwoZeros", () => {
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

  describe("maybeToString", () => {
    it("should return string representation of truthy values", () => {
      expect(maybeToString(123)).toBe("123");
      expect(maybeToString("hello")).toBe("hello");
      expect(maybeToString(true)).toBe("true");
    });

    it("should return empty string for falsy values", () => {
      expect(maybeToString(0)).toBe("");
      expect(maybeToString(null)).toBe("");
      expect(maybeToString(undefined)).toBe("");
      expect(maybeToString(false)).toBe("");
      expect(maybeToString("")).toBe("");
    });
  });
});
