import { Flag } from "./components/icons/Flag";
import { Mine } from "./components/icons/Mine";
import { Number } from "./components/icons/Number";
import { Square } from "./components/icons/Square";

function App() {
  return (
    <main className="p-2">
      <h1 className="text-xl">mines</h1>
      <div className="flex flex-row">
        <Number val={2} />
        <Square />
        <Flag />
        <Mine />
      </div>
      <div className="flex flex-row">
        <Square />
        <Number val={0} />
        <Number val={2} />
        <Mine />
      </div>
    </main>
  );
}

export default App;
