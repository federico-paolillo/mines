import { render } from "preact";
import "preact/debug";
import { App } from "./App";

const appNode = document.createElement("div");

document.body.appendChild(appNode);

render(<App />, appNode);
