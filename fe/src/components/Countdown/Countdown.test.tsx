import { render, screen, act } from "@testing-library/preact";
import { describe, expect, it, vi, beforeEach, afterEach } from "vitest";
import { Countdown } from "./index";

describe("Countdown", () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  it("should display initial time correctly", () => {
    const startTime = Math.floor(Date.now() / 1000);
    const onExpired = vi.fn();
    render(<Countdown startTime={startTime} onExpired={onExpired} />);

    // 2 hours = 02:00:00
    expect(screen.getByTestId("countdown")).toHaveTextContent("02:00:00");
  });

  it("should countdown correctly", () => {
    const startTime = Math.floor(Date.now() / 1000);
    const onExpired = vi.fn();
    render(<Countdown startTime={startTime} onExpired={onExpired} />);

    act(() => {
      vi.advanceTimersByTime(1000);
    });
    expect(screen.getByTestId("countdown")).toHaveTextContent("01:59:59");

    act(() => {
      vi.advanceTimersByTime(1000);
    });
    expect(screen.getByTestId("countdown")).toHaveTextContent("01:59:58");
  });

  it("should call onExpired when time is up", () => {
    const startTime = Math.floor(Date.now() / 1000);
    const onExpired = vi.fn();
    render(<Countdown startTime={startTime} onExpired={onExpired} />);

    // Advance 2 hours + 1 second
    act(() => {
      vi.advanceTimersByTime(2 * 60 * 60 * 1000 + 1000);
    });

    expect(screen.getByTestId("countdown")).toHaveTextContent("00:00:00");
    expect(onExpired).toHaveBeenCalled();
  });

  it("should handle already expired start time", () => {
    const startTime = Math.floor(Date.now() / 1000) - (2 * 60 * 60 + 100); // Expired 100s ago
    const onExpired = vi.fn();
    render(<Countdown startTime={startTime} onExpired={onExpired} />);

    expect(screen.getByTestId("countdown")).toHaveTextContent("00:00:00");
    expect(onExpired).toHaveBeenCalled();
  });
});
