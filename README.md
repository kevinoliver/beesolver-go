# Spelling Bee Solver

Solves the [New York Times Spelling Bee](https://www.nytimes.com/puzzles/spelling-bee).

Their dictionary is not the same as the one included so you find missing words as well as extras.

## Usage

```
$ go build *.go
$ ./beesolver --help
Usage: beesolver REQUIRED_LETTER OTHER_LETTERS
      -dict string
            Path to a custom dictionary
      -help
            Print the usage and exit
      -words-output
            Default on. When off, the solution's words are hidden (default true)
```

## Requirements

* Go 1.20 — though given how little language functionality is used it likely works on much earlier versions of Go.

## Why

This was an excuse to play around with Go. I come from a Java background and
[did this in Java first](https://github.com/kevinoliver/beesolver-java). Unsurprisingly,
until I learn the language's idioms and conventions, this code feels quite Java-y.

## Development

Run all tests:
```
$ go test -v ./...
```