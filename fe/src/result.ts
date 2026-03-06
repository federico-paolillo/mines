export interface Problem {
  message: string;
  cause?: Error;
}

export interface Success<T> {
  success: true;
  value: T;
}

export interface Failure<E = Problem> {
  success: false;
  error: E;
}

export type Result<T, E = Problem> = Success<T> | Failure<E>;

export type ErrorKind = "match_over" | "unknown";

export interface ApiError {
  kind: ErrorKind;
  message: string;
}

export function success<T>(value: T): Result<T> {
  return {
    success: true,
    value,
  };
}

export function failure<T = never, E = Problem>(error: E): Result<T, E> {
  return {
    success: false,
    error,
  };
}
