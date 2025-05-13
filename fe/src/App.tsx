import { Board } from "./components/Board";

function App() {
  return (
    <main className="p-2">
      <h1 className="text-xl">mines</h1>
      <Board
        width={2}
        height={2}
        cells={[
          { x: 0, y: 0, state: "closed" },
          { x: 1, y: 0, state: "closed" },
          { x: 0, y: 0, state: "closed" },
          { x: 1, y: 1, state: "closed" },
        ]}
      />
    </main>
  );
}

export default App;
