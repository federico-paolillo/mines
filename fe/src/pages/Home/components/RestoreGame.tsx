import { useState } from "preact/hooks";
import { useLocation } from "preact-iso";
import type { FunctionalComponent } from "preact";

const RestoreGame: FunctionalComponent = () => {
  const [gameId, setGameId] = useState("");
  const { route } = useLocation();

  const isButtonDisabled = gameId.length < 8;

  const handleSubmit = (e: Event) => {
    e.preventDefault();
    if (!isButtonDisabled) {
      route(`/game?id=${gameId}`);
    }
  };

  const handleInput = (e: Event) => {
    const target = e.target as HTMLInputElement;
    setGameId(target.value);
  };

  return (
    <div class="mt-8">
      <form class="flex items-center" onSubmit={handleSubmit}>
        <input
          type="text"
          required
          value={gameId}
          onInput={handleInput}
          class="bg-gray-200 appearance-none border-2 border-gray-200 rounded py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
          placeholder="Enter Game ID"
        />
        <button
          class="ml-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded disabled:bg-gray-400"
          type="submit"
          disabled={isButtonDisabled}
        >
          Restore Game
        </button>
      </form>
    </div>
  );
};

export { RestoreGame };
