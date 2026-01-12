import { act, render, screen } from "@testing-library/preact";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { toUnixTimestamp } from "../../time";
import { Countdown } from "./index";

describe("Countdown", () => {
  const DURATION_SECONDS = 2 * 60 * 60; // 2 hours

  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  it("should display initial time correctly", () => {
    const startTime = toUnixTimestamp(new Date());
    const onExpired = vi.fn();
    render(
      <Countdown
        startTime={startTime}
        durationSeconds={DURATION_SECONDS}
        onExpired={onExpired}
      />,
    );

    // 2 hours = 02:00:00
    expect(screen.getByTestId("countdown").textContent).toBe("02:00:00");
  });

  it("should countdown correctly", () => {
    const startTime = toUnixTimestamp(new Date());
    const onExpired = vi.fn();
    render(
      <Countdown
        startTime={startTime}
        durationSeconds={DURATION_SECONDS}
        onExpired={onExpired}
      />,
    );

    act(() => {
      vi.advanceTimersByTime(1000);
    });
    expect(screen.getByTestId("countdown").textContent).toBe("01:59:59");

    act(() => {
      vi.advanceTimersByTime(1000);
    });
    expect(screen.getByTestId("countdown").textContent).toBe("01:59:58");
  });

  it("should call onExpired when time is up", () => {
    const startTime = toUnixTimestamp(new Date());
    const onExpired = vi.fn();
    render(
      <Countdown
        startTime={startTime}
        durationSeconds={DURATION_SECONDS}
        onExpired={onExpired}
      />,
    );

    // Advance 2 hours + 1 second
    act(() => {
      vi.advanceTimersByTime(DURATION_SECONDS * 1000 + 1000);
    });

    expect(screen.getByTestId("countdown").textContent).toBe("00:00:00");
    expect(onExpired).toHaveBeenCalled();
  });

  it("should handle already expired start time", () => {
    const startTime = toUnixTimestamp(new Date()) - (DURATION_SECONDS + 100); // Expired 100s ago
    const onExpired = vi.fn();
    render(
      <Countdown
        startTime={startTime}
        durationSeconds={DURATION_SECONDS}
        onExpired={onExpired}
      />,
    );

    expect(screen.getByTestId("countdown").textContent).toBe("00:00:00");
    expect(onExpired).toHaveBeenCalled();
  });
});
