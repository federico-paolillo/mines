import { useLocation } from "preact-iso";
import { DifficultyObject } from "../../client/models/game";
import { useApiClient } from "../../clientContext";
import { NewGame } from "./components/NewGame";
import { RestoreGame } from "./components/RestoreGame";

export function Home() {
  const { route } = useLocation();
  const client = useApiClient();

  const handleNewGame = async () => {
    const result = await client.startNewGame({
      difficulty: DifficultyObject.Beginner,
    });

    if (result.success) {
      const matchId = result.value.id;
      // Navigate to the game page with the new game id
      route(`/game?id=${matchId}`);
    } else {
      console.error(result.error);
    }
  };

  return (
    <div class="container mx-auto p-4 text-center">
      <h1 class="text-4xl font-bold">Welcome to Mines</h1>
      <p class="mt-4 text-lg">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.
        Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies
        sed, dolor. Cras elementum ultrices diam. Maecenas ligula massa, varius
        a, semper congue, euismod non, mi.
      </p>

      <div class="mt-8 grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="p-6 border rounded-lg">
          <h2 class="text-2xl font-bold">Start a New Game</h2>
          <p class="mt-2">
            Click the button below to start a fresh game of Minesweeper. The
            board will be generated for you.
          </p>
          <NewGame onNewGame={handleNewGame} />
        </div>

        <div class="p-6 border rounded-lg">
          <h2 class="text-2xl font-bold">Load an Existing Game</h2>
          <p class="mt-2">
            If you have a game ID, you can restore your previous game session by
            entering it below.
          </p>
          <RestoreGame />
        </div>
      </div>
    </div>
  );
}
