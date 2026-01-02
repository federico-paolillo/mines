import type { FunctionalComponent } from "preact";

const NewGame: FunctionalComponent = () => {
  return (
    <div class="mt-8">
      <button
        type="button"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        New Game
      </button>
    </div>
  );
};

export { NewGame };
