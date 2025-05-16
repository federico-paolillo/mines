export interface SpinnerProps {
  isLoading: true;
  children: React.ReactNode;
}

export function Spinner({ isLoading, children }: SpinnerProps) {
  return (
    <div className="relative w-[100%] h-[100%]">
      {children}
      {isLoading ? (
        <div className="absolute top-0 left-0 w-[100%] h-[100%] bg-[#bbb9] flex justify-center items-center">
          <span className="text-2xl text-black z-10 animate-spin">꩜</span>
        </div>
      ) : null}
    </div>
  );
}
