export interface SpinnerProps {
  isOpen: boolean;
}

export function Spinner({ isOpen }: SpinnerProps) {
  if (!isOpen) {
    return null;
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
      aria-label="Loading"
      role="progressbar"
    >
      <div className="w-16 h-16 border-4 border-gray-200 border-t-gray-800 rounded-full animate-spin" />
    </div>
  );
}
