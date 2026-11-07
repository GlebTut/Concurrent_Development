# Lab 4 - Barrier Abstract Data Type

A reusable barrier as an abstract data type: a `barrier` struct created with `createBarrier(N)` and used through its `wait()` method. It uses an atomic counter and two unbuffered channels (two turnstiles), the same design as SimpleBarrier. Five goroutines use it for 3 rounds.

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod`). Only the standard library is used.

## How to run

From the root of the repository:

```
go run ./lab_four/BarrierADT
```

## Expected output

For each round, five `Part A round <r> goroutine <n>` lines appear, then five `PartB round <r> goroutine <n>` lines. The order of the goroutine numbers changes from run to run, but no `PartB` line of a round appears before all five `Part A` lines of that round.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
