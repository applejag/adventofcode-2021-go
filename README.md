# Advent of Code 2021 via Go

This repo contains my attempt at Advent of Code 2021
(<https://adventofcode.com/2021>).

## Running it

- Install Go 1.17 (or higher)

- Run the `go run` CLI inside the package you wish to run. For day 1, you
  would run:

  ```console
  $ cd cmd/day01
  $ go run .
  [INFO |common|…common.go:72] Reading file.  path=input.txt
  [INFO |day01 |…/day01.go:44] Scanning complete.  scans=2000  increases=1791
  ```

  To run part 2 of the puzzle, add the `-2` flag:

  ```console
  $ cd cmd/day01
  $ go run . -2
  [INFO |common|…common.go:72] Reading file.  path=input.txt
  [INFO |day01 |…/day01.go:22] Using windowed scanner.  window=3
  [INFO |day01 |…/day01.go:44] Scanning complete.  scans=1998  increases=1822
  ```
