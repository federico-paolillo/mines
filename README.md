# mines

[![Go Report Card](https://goreportcard.com/badge/github.com/federico-paolillo/mines)](https://goreportcard.com/report/github.com/federico-paolillo/mines)
[![Build status](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml/badge.svg)](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/federico-paolillo/mines/branch/main/graph/badge.svg?token=N8BYYY510Z)](https://codecov.io/github/federico-paolillo/mines)

## What is this ?

> Minesweeper is a logic puzzle video game genre generally played on personal computers. The game features a grid of clickable tiles, with hidden "mines" (depicted as naval mines in the original game) scattered throughout the board. The objective is to clear the board without detonating any mines, with help from clues about the number of neighboring mines in each field.

**From:** https://en.wikipedia.org/wiki/Minesweeper_(video_game)

It's Minesweeper in Go.

This project has a server that let's you play matches of Minesweeper through an HTTP REST-like API. You will be able to request a match and play them until you win, lose or 2 hours pass and the game is automatically destroyed. The server component is written in Go using Gin, while the client component is a Single Page Application written in React using Vite as the bundler of choice.
