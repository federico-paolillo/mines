# mines

[![Go Report Card](https://goreportcard.com/badge/github.com/federico-paolillo/mines)](https://goreportcard.com/report/github.com/federico-paolillo/mines)
[![Build status](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml/badge.svg)](https://github.com/federico-paolillo/mines/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/federico-paolillo/mines/branch/main/graph/badge.svg?token=N8BYYY510Z)](https://codecov.io/github/federico-paolillo/mines)

**Note:** This software is in development, not everything will work.

[Minesweeper](<https://en.wikipedia.org/wiki/Minesweeper_(video_game)>). Minesweeper is a logic puzzle video game genre generally played on personal computers. The game features a grid of clickable tiles, with hidden "mines" (depicted as naval mines in the original game) scattered throughout the board. The objective is to clear the board without detonating any mines, with help from clues about the number of neighboring mines in each field.

This implementation of Minesweeper is a tool to investigate "out-of-process" frontends (as in [Aptitude](https://wiki.debian.org/Aptitude) **is a frontend** for [Apt](https://wiki.debian.org/Apt) libraries). By default, this game is somewhat playable via a remarkably (and purposefully) terrible built-in TUI. The program can be instructed to run in **server mode** to run a server listening on an Unix-socket (no Microsoft Windows, sorry). The server will accept and apply game commands and will return the updated game state. Hopefully, this should enable building frontends with any technology.

This software is meant and designed to work on Unix-like systems only. I will not make try to make it compatibile with Microsoft Windows. For the moment, there are no third-party dependencies save for the Go standard library and Cobra.
