import type { FunctionalComponent } from "preact";

interface NewGameProps {
  onNewGame: () => void;
}

const NewGame: FunctionalComponent<NewGameProps> = ({ onNewGame }) => {
  const handleClick = () => {
    onNewGame();
  };

  return (
    <div class="mt-8">
      <button
        type="button"
        onClick={handleClick}
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        New Game
      </button>
    </div>
  );
};

export { NewGame };
