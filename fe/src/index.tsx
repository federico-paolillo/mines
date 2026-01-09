import { render } from "preact";
import { LocationProvider, Route, Router } from "preact-iso";
import "preact/debug";

import { ClientProvider } from "./clientContext.tsx";
import { Header } from "./components/Header.tsx";
import { NotFound } from "./pages/_404.tsx";
import { Game } from "./pages/Game/index.tsx";
import { Home } from "./pages/Home/index.tsx";
import "./style.css";

export function App() {
  return (
    <LocationProvider>
      <ClientProvider>
        <Header />
        <main>
          <Router>
            <Route path="/" component={Home} />
            <Route path="/game" component={Game} />
            <Route default component={NotFound} />
          </Router>
        </main>
      </ClientProvider>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
