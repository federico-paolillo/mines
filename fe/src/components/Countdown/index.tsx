import { useEffect, useState } from "preact/hooks";
import { formatSecondsToHhMmSs, toUnixTimestamp } from "../../time";

interface CountdownProps {
  startTime: number; // Unix timestamp in seconds
  durationSeconds: number; // Duration in seconds
  onExpired: () => void;
}

export function Countdown({ startTime, durationSeconds, onExpired }: CountdownProps) {
  // startTime is in seconds, convert to milliseconds for date-fns if needed
  // or just work with seconds since we are counting down seconds.

  // We want to count down from (startTime + duration) - now.

  const [timeLeft, setTimeLeft] = useState<number>(0);

  useEffect(() => {
    const calculateTimeLeft = () => {
      const now = toUnixTimestamp(new Date()); // Current time in seconds
      const endTime = startTime + durationSeconds;
      const remaining = endTime - now;

      if (remaining <= 0) {
        setTimeLeft(0);
        onExpired();
        return 0;
      }
      return remaining;
    };

    // Initial calculation
    const initialRemaining = calculateTimeLeft();
    setTimeLeft(initialRemaining);

    // If expired immediately, no need for interval (handled in calculateTimeLeft)
    if (initialRemaining <= 0) return;

    const intervalId = setInterval(() => {
      const remaining = calculateTimeLeft();
      if (remaining <= 0) {
        clearInterval(intervalId);
      } else {
        setTimeLeft(remaining);
      }
    }, 1000);

    return () => clearInterval(intervalId);
  }, [startTime, durationSeconds, onExpired]);

  return (
    <div data-testid="countdown">
      {formatSecondsToHhMmSs(timeLeft)}
    </div>
  );
}
