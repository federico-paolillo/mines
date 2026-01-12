import { intervalToDuration } from "date-fns";
import { padWithTwoZeros } from "./strings";

/**
 * Converts a Date object to a UNIX timestamp in seconds.
 * @param date The date to convert.
 * @returns The UNIX timestamp in seconds.
 */
export function toUnixTimestamp(date: Date): number {
  return Math.floor(date.getTime() / 1000);
}

/**
 * Formats a duration in seconds to HH:MM:SS string.
 * @param seconds The duration in seconds.
 * @returns The formatted time string.
 */
export function formatSecondsToHhMmSs(seconds: number): string {
  const duration = intervalToDuration({ start: 0, end: seconds * 1000 });

  // intervalToDuration returns years, months, days, hours, minutes, seconds
  const h = duration.hours || 0;
  const m = duration.minutes || 0;
  const s = duration.seconds || 0;

  return `${padWithTwoZeros(h.toString())}:${padWithTwoZeros(m.toString())}:${padWithTwoZeros(s.toString())}`;
}

/**
 * Calculates the remaining time in seconds.
 * @param startTimeInSeconds The start time in UNIX seconds.
 * @param durationInSeconds The duration in seconds.
 * @param nowInSeconds The current time in UNIX seconds.
 * @returns The remaining time in seconds, or 0 if expired.
 */
export function calculateTimeLeftInSeconds(
  startTimeInSeconds: number,
  durationInSeconds: number,
  nowInSeconds: number,
): number {
  const endTimeInSeconds = startTimeInSeconds + durationInSeconds;
  const remainingTimeInSeconds = endTimeInSeconds - nowInSeconds;
  return Math.max(0, remainingTimeInSeconds);
}
