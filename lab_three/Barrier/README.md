# Lab 3 - Barrier

A reusable N-goroutine barrier in Go, built from a mutex and two weighted semaphores (a two-turnstile design).

Ten goroutines each do "Part A" (after a one second sleep), then wait at the barrier.
Only when all ten have finished Part A does any of them start "Part B".

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod` of the repository root)
- Internet access the first time, so Go can download `golang.org/x/sync`

## How to run

From the root of the repository:

```
go mod tidy
go run ./lab_three/Barrier
```

Or from inside this folder:

```
go run barrier.go
```

## Expected output

After about one second, ten `Part A <n>` lines appear, followed by ten `Part B <n>` lines.
The order of the numbers inside each group changes from run to run, but no `Part B` line
ever appears before every `Part A` line has been printed.

## Tests

The tests check that no goroutine passes the barrier before all of them have arrived, and that the same barrier can be used for 100 rounds in a row without a deadlock or a goroutine running ahead.

From the root of the repository:

```
go test ./lab_three/Barrier
```

Add `-race` (needs cgo and gcc) to also look for data races.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
