# mines

[![Go Report Card](https://goreportcard.com/badge/github.com/federico-paolillo/mines)](https://goreportcard.com/report/github.com/federico-paolillo/mines)
[![Build status](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml/badge.svg)](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/federico-paolillo/mines/branch/main/graph/badge.svg?token=N8BYYY510Z)](https://codecov.io/github/federico-paolillo/mines)

**Note:** This software is in development. It may be that some of the things I write here do not yet exist.

[Minesweeper](<https://en.wikipedia.org/wiki/Minesweeper_(video_game)>). Minesweeper is a logic puzzle video game genre generally played on personal computers. The game features a grid of clickable tiles, with hidden "mines" (depicted as naval mines in the original game) scattered throughout the board. The objective is to clear the board without detonating any mines, with help from clues about the number of neighboring mines in each field.

This implementation of Minesweeper has the goal of investigating pluggable frontends (as in [Aptitude](https://wiki.debian.org/Aptitude) **is a frontend** for [Apt](https://wiki.debian.org/Apt) libraries). By default, this game is playable via a built-in TUI but the program can be instructed to run in **backend mode**. In backend mode the program will expose an Unix-socket server (no Microsoft Windows, sorry) that will accept game commands and will return new game state after applying said commands. This should hopefully make possible to build any frontend with any technology.

Eventually, shall I find the time, another Desktop-only frontend will be provided using [Tauri](https://tauri.app/).

## Usage

You need [Go](https://go.dev/dl/) 1.22.1. This software is meant and designed to work on Unix-like systems. I will not make any attempts at compatibility with Microsoft Windows. For the moment, this software has no third-party dependencies save the Go standard library.

From the source root folder, run `go test ./...` to run tests or run `go run ./cmd/cli` to run the built-in client.
