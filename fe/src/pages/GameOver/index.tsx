export function GameOver() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
      <h1 className="text-4xl font-bold mb-4 text-red-600">Game Over</h1>
      <p className="text-xl mb-8">The match has expired or you ran out of lives.</p>
      <a href="/" className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
        Return to Home
      </a>
    </div>
  );
}
