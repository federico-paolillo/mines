# mines

[![Go Report Card](https://goreportcard.com/badge/github.com/federico-paolillo/mines)](https://goreportcard.com/report/github.com/federico-paolillo/mines)
[![Build status](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml/badge.svg)](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/federico-paolillo/mines/branch/main/graph/badge.svg?token=N8BYYY510Z)](https://codecov.io/github/federico-paolillo/mines)

## What is this ?

> Minesweeper is a logic puzzle video game genre generally played on personal computers. The game features a grid of clickable tiles, with hidden "mines" (depicted as naval mines in the original game) scattered throughout the board. The objective is to clear the board without detonating any mines, with help from clues about the number of neighboring mines in each field.

**From:** https://en.wikipedia.org/wiki/Minesweeper_(video_game)

It's Minesweeper in Go.

This project has a server that let's you play matches of Minesweeper through an HTTP REST-like API. You will be able to request a match and play them until you win, lose or 2 hours pass and the game is automatically destroyed. The server component is written in Go using Gin, while the client component is a Single Page Application written in React using Vite as the bundler of choice.

## Design Decisions

### useGameState

The `useGameState` hook encapsulates the game state management and interaction logic. I decided to include the data fetching responsibility within this hook.

While separating data fetching from interaction logic can sometimes provide better separation of concerns (e.g., using a dedicated data fetching hook and a separate logic hook), in this specific case, the game state is tightly coupled with the interactions. Every interaction (click) results in a server call that updates the game state.

By bundling fetching and interactions together:
1.  **Simplicity**: The consuming component (`Game`) becomes truly "dumb", only responsible for rendering the state it receives. It doesn't need to coordinate between a fetching hook and a logic hook.
2.  **Consistency**: The hook ensures that the `gameState` returned is always in sync with the latest server response, whether it comes from the initial fetch or a subsequent move.
3.  **Encapsulation**: All API error handling, redirect logic (e.g., game over), and loading states are contained in one place, reducing code duplication and potential bugs in the consumer.

If the application were to grow significantly, or if we needed to reuse the fetching logic in a context where interactions are not needed (e.g., a read-only spectator mode), we might reconsider this decision and split the responsibilities.
