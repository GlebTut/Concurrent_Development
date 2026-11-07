# Lab 4 - Simple Barrier

A reusable barrier for 10 goroutines, built from an atomic counter and two unbuffered channels (two turnstiles). Each goroutine runs 3 rounds of "Part A", waits at the barrier, then does "Part B".

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod`). Only the standard library is used.

## How to run

From the root of the repository:

```
go run ./lab_four/SimpleBarrier
```

## Expected output

For each round, ten `Part A round <r> goroutine <n>` lines appear, then ten `PartB round <r> goroutine <n>` lines. The order of the goroutine numbers changes from run to run, but no `PartB` line of a round appears before all ten `Part A` lines of that round.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
