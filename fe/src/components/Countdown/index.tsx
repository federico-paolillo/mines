import { useEffect, useState } from "preact/hooks";
import {
  calculateTimeLeftInSeconds,
  formatSecondsToHhMmSs,
  toUnixTimestamp,
} from "../../time";

interface CountdownProps {
  startTime: number; // Unix timestamp in seconds
  durationSeconds: number; // Duration in seconds
  onExpired: () => void;
}

export function Countdown({
  startTime,
  durationSeconds,
  onExpired,
}: CountdownProps) {
  // Initialize with current calculation to avoid flash of 00:00:00
  const [timeLeft, setTimeLeft] = useState<number>(() =>
    calculateTimeLeftInSeconds(
      startTime,
      durationSeconds,
      toUnixTimestamp(new Date()),
    ),
  );

  useEffect(() => {
    // Check immediately on mount/update
    const currentRemaining = calculateTimeLeftInSeconds(
      startTime,
      durationSeconds,
      toUnixTimestamp(new Date()),
    );
    setTimeLeft(currentRemaining);

    if (currentRemaining === 0) {
      onExpired();
      return;
    }

    const intervalHandle = setInterval(() => {
      const remaining = calculateTimeLeftInSeconds(
        startTime,
        durationSeconds,
        toUnixTimestamp(new Date()),
      );
      setTimeLeft(remaining);

      if (remaining === 0) {
        clearInterval(intervalHandle);
        onExpired();
      }
    }, 1000);

    return () => clearInterval(intervalHandle);
  }, [startTime, durationSeconds, onExpired]);

  return <div data-testid="countdown">{formatSecondsToHhMmSs(timeLeft)}</div>;
}
