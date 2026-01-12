/**
 * Converts a Date object to a UNIX timestamp in seconds.
 * @param date The date to convert.
 * @returns The UNIX timestamp in seconds.
 */
export function toUnixTimestamp(date: Date): number {
  return Math.floor(date.getTime() / 1000);
}
