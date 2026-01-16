interface LifeContainerProps {
  lives: number;
}

export function LifeContainer({ lives }: LifeContainerProps) {
  const hearts = Array.from({ length: lives }, (_, index) => index);

  return (
    <div className="flex items-center gap-1" role="status" aria-label={`${lives} lives remaining`}>
      {hearts.map((index) => (
        <span
          key={index}
          className="text-red-500 text-2xl"
          aria-hidden="true"
        >
          ♥
        </span>
      ))}
    </div>
  );
}
