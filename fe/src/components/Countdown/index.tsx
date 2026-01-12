import { intervalToDuration } from "date-fns";
import { useEffect, useState } from "preact/hooks";
import { toUnixTimestamp } from "../../time";

interface CountdownProps {
  startTime: number; // Unix timestamp in seconds
  onExpired: () => void;
}

const MATCH_DURATION_SECONDS = 2 * 60 * 60; // 2 hours

export function Countdown({ startTime, onExpired }: CountdownProps) {
  // startTime is in seconds, convert to milliseconds for date-fns if needed
  // or just work with seconds since we are counting down seconds.

  // We want to count down from (startTime + 2h) - now.

  const [timeLeft, setTimeLeft] = useState<number>(0);

  useEffect(() => {
    const calculateTimeLeft = () => {
      const now = toUnixTimestamp(new Date()); // Current time in seconds
      const endTime = startTime + MATCH_DURATION_SECONDS;
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
  }, [startTime, onExpired]);

  const formatTime = (seconds: number) => {
    const duration = intervalToDuration({ start: 0, end: seconds * 1000 });

    // intervalToDuration returns years, months, days, hours, minutes, seconds
    // Since we only care about HH:MM:SS and the total is max 2 hours:
    const h = duration.hours || 0;
    const m = duration.minutes || 0;
    const s = duration.seconds || 0;

    const pad = (n: number) => n.toString().padStart(2, "0");
    return `${pad(h)}:${pad(m)}:${pad(s)}`;
  };

  return (
    <div data-testid="countdown">
      {formatTime(timeLeft)}
    </div>
  );
}
