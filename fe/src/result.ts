export type ErrorKind = "match_over" | "unknown";

export interface Problem {
  kind: ErrorKind;
  message: string;
  cause?: Error;
}

export interface Success<T> {
  success: true;
  value: T;
}

export interface Failure {
  success: false;
  error: Problem;
}

export type Result<T> = Success<T> | Failure;

export function success<T>(value: T): Result<T> {
  return {
    success: true,
    value,
  };
}

export function failure<T = never>(problem: Problem): Result<T> {
  return {
    success: false,
    error: problem,
  };
}

export function isMatchOver(result: Failure): boolean {
  return result.error.kind === "match_over";
}
