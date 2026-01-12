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
