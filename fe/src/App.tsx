import { LocationProvider, Route, Router } from "preact-iso";
import "preact/debug";

import { ClientProvider } from "./clientContext.tsx";
import { Header } from "./components/Header.tsx";
import { NotFound } from "./pages/_404.tsx";
import { Game } from "./pages/Game/index.tsx";
import { Home } from "./pages/Home/index.tsx";

import "./style.css";

import { config } from "./config.ts";

export function App() {
  return (
    <LocationProvider>
      <ClientProvider apiBaseUrl={config.apiBaseUrl}>
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
