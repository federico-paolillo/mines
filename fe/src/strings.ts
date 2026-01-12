/**
 * Pads a string with leading zeros to ensure it has at least length 2.
 * @param input The string to pad.
 * @returns The padded string.
 */
export function padWithTwoZeros(input: string): string {
  return input.padStart(2, "0");
}
