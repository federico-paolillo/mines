import { Board } from "./components/Board";
import { Spinner } from "./components/Spinner";

function App() {
  return (
    <main className="p-2">
      <h1 className="text-xl">mines</h1>
      <div className="w-fit">
        <Spinner isLoading={true}>
          <Board
            width={3}
            height={3}
            cells={[
              { x: 0, y: 0, state: "open" },
              { x: 1, y: 0, state: "flagged" },
              { x: 2, y: 0, state: "closed" },
              { x: 0, y: 1, state: "unfathomable" },
              { x: 1, y: 1, state: "flagged" },
              { x: 2, y: 1, state: "closed" },
              { x: 0, y: 2, state: "closed" },
              { x: 1, y: 2, state: "closed" },
              { x: 2, y: 2, state: "closed" },
            ]}
          />
        </Spinner>
      </div>
    </main>
  );
}

export default App;
