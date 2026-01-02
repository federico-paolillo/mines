import { useLocation, useRoute } from "preact-iso";

export function Game() {
  const { query } = useRoute();
  const params = new URLSearchParams(query);
  const gameId = params.get("id");

  return (
    <div class="container mx-auto p-4">
      <h1 class="text-2xl font-bold">Game Page</h1>
      {gameId ? (
        <p>Loading game with ID: {gameId}</p>
      ) : (
        <p>Starting a new game...</p>
      )}
    </div>
  );
}
