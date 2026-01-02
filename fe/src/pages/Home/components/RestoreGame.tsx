import type { FunctionalComponent } from "preact";

const RestoreGame: FunctionalComponent = () => {
  return (
    <div class="mt-8">
      <form class="flex items-center">
        <input
          type="text"
          class="bg-gray-200 appearance-none border-2 border-gray-200 rounded py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
          placeholder="Enter Game ID"
        />
        <button
          class="ml-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
          type="submit"
        >
          Restore Game
        </button>
      </form>
    </div>
  );
};

export { RestoreGame };
