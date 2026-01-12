import { describe, expect, it } from "vitest";
import {
  calculateTimeLeftInSeconds,
  formatSecondsToHhMmSs,
  toUnixTimestamp,
} from "./time";

describe("time", () => {
  describe("toUnixTimestamp", () => {
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

  describe("formatSecondsToHhMmSs", () => {
    it("should format 0 seconds", () => {
      expect(formatSecondsToHhMmSs(0)).toBe("00:00:00");
    });

    it("should format seconds less than a minute", () => {
      expect(formatSecondsToHhMmSs(45)).toBe("00:00:45");
    });

    it("should format minutes and seconds", () => {
      expect(formatSecondsToHhMmSs(125)).toBe("00:02:05");
    });

    it("should format hours, minutes and seconds", () => {
      expect(formatSecondsToHhMmSs(3665)).toBe("01:01:05");
    });
  });

  describe("calculateTimeLeftInSeconds", () => {
    it("should calculate remaining time correctly", () => {
      const start = 1000;
      const duration = 100;
      const now = 1050;
      expect(calculateTimeLeftInSeconds(start, duration, now)).toBe(50);
    });

    it("should return 0 if time is up", () => {
      const start = 1000;
      const duration = 100;
      const now = 1100;
      expect(calculateTimeLeftInSeconds(start, duration, now)).toBe(0);
    });

    it("should return 0 if time is passed", () => {
      const start = 1000;
      const duration = 100;
      const now = 1200;
      expect(calculateTimeLeftInSeconds(start, duration, now)).toBe(0);
    });

    it("should return duration if now is start time", () => {
      const start = 1000;
      const duration = 100;
      const now = 1000;
      expect(calculateTimeLeftInSeconds(start, duration, now)).toBe(100);
    });
  });
});
