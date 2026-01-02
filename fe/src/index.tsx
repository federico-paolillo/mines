import "preact/debug";
import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";

import { Header } from "./components/Header.tsx";
import { Game } from "./pages/Game/index.tsx";
import { Home } from "./pages/Home/index.tsx";
import { NotFound } from "./pages/_404.tsx";
import "./style.css";

export function App() {
  return (
    <LocationProvider>
      <Header />
      <main>
        <Router>
          <Route path="/" component={Home} />
          <Route path="/game" component={Game} />
          <Route default component={NotFound} />
        </Router>
      </main>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
