export function Flag() {
  return (
    <div className="cursor-pointer w-6 h-6 flex items-center justify-center text-xs select-none bg-[#c0c0c0] hover:bg-[#b0b0b0] active:bg-[#a0a0a0] border-[3px] border-t-gray-100 border-l-gray-100 border-b-[#86888c] border-r-[#86888c]">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <line x1="12" y1="5" x2="12" y2="19" stroke="black" stroke-width="3" />
        <line x1="8" y1="19" x2="16" y2="19" stroke="black" stroke-width="3" />
        <polygon points="12,5 20,8 12,11" fill="red" stroke="none" />
      </svg>
    </div>
  );
}
