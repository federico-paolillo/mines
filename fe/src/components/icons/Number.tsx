import { twJoin } from "tailwind-merge";

export interface NumberProps {
  val: number;
}

const DEFAULT_CELL_COLOR = "text-black";

const valToColorMap = new Map<number, string>([
  [1, "text-blue-600"],
  [2, "text-green-700"],
  [3, "text-red-600"],
  [4, "text-blue-900"],
  [5, "text-red-800"],
  [6, "text-teal-600"],
  [7, "text-black"],
  [8, "text-gray-600"],
]);

function formatNeighboursText(val: number): string {
  if (val == 0) {
    return "";
  }

  return String(val);
}

export function Number({ val }: NumberProps) {
  const neighboursText = formatNeighboursText(val);
  const neighboursTextColor = valToColorMap.get(val) ?? DEFAULT_CELL_COLOR;

  return (
    <div className="w-6 h-6 flex items-center justify-center text-xs select-none bg-[#c0c0c0] border border-[#a3a8ac]">
      <span className={twJoin("text-bold text-xl", neighboursTextColor)}>
        {neighboursText}
      </span>
    </div>
  );
}
