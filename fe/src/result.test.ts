import { describe, expect, it } from "vitest";
import { failure, type Problem, type Result, success } from "./result";

describe("Result Pattern", () => {
  it("should create a success result with a value", () => {
    const value = 42;
    const result = success(value);

    expect(result.success).toBe(true);

    if (result.success) {
      expect(result.value).toBe(value);
    }
  });

  it("should create a failure result with a problem", () => {
    const problem: Problem = { message: "Something went wrong" };
    const result = failure(problem);

    expect(result.success).toBe(false);

    if (!result.success) {
      expect(result.error).not.toBeNull();
      expect(result.error).toEqual(problem);
    }
  });

  it("should support optional cause in Problem", () => {
    const cause = new Error("Root cause");
    const problem: Problem = { message: "Wrapped error", cause };
    const result = failure(problem);

    if (!result.success) {
      expect(result.error.cause).toBe(cause);
      expect(result.error.message).toBe("Wrapped error");
    }
  });

  it("should allow type narrowing via discriminated union", () => {
    const processResult = (res: Result<string>): string => {
      if (res.success) {
        return `Success: ${res.value}`;
      }
      return `Error: ${res.error.message}`;
    };

    expect(processResult(success("ok"))).toBe("Success: ok");
    expect(processResult(failure({ message: "fail" }))).toBe("Error: fail");
  });
});
