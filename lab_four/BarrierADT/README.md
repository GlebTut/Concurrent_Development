# Lab 4 - Barrier Abstract Data Type

A reusable barrier as an abstract data type: a `barrier` struct created with `createBarrier(N)` and used through its `wait()` method. It uses an atomic counter and two unbuffered channels (two turnstiles), the same design as SimpleBarrier. Five goroutines use it for 3 rounds.

## GitHub

All the labs are in this repository: https://github.com/GlebTut/Concurrent_Development

This lab: https://github.com/GlebTut/Concurrent_Development/tree/main/lab_four/BarrierADT

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## How it works

This is the same two-phase barrier as in `SimpleBarrier`, but packaged as an abstract data type: the
`barrier` struct holds everything the barrier needs, and the goroutines only call `wait()`.

- `createBarrier(N)` returns a pointer to a new barrier for `N` goroutines. It returns a pointer so
  the barrier (which contains an atomic counter) is never copied.
- `count` is an atomic counter of the goroutines inside the barrier. `turnstile1` and `turnstile2`
  are unbuffered channels used as signals.
- `wait()` **phase 1:** add 1 to `count`. The last goroutine to arrive sends `total-1` signals on
  `turnstile1`; the others receive one and continue.
- `wait()` **phase 2:** subtract 1 from `count`. The last goroutine to leave sends `total-1` signals
  on `turnstile2`. Nobody can enter the barrier again until everybody has left it, so `wait()` can
  be called in every round.

The template used one channel and a mutex, and it could only be used once. Here the mutex was
replaced by the atomic counter and the second turnstile was added to make it reusable.

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod`). Only the standard library is used.

## How to run

From the root of the repository:

```
go run ./lab_four/BarrierADT
```

## Expected output

For each round, five `Part A round <r> goroutine <n>` lines appear, then five `PartB round <r> goroutine <n>` lines. The order of the goroutine numbers changes from run to run, but no `PartB` line of a round appears before all five `Part A` lines of that round.

## Tests

The tests check that no goroutine passes the barrier before all of them have arrived, and that the same barrier can be used for 100 rounds in a row without a deadlock or a goroutine running ahead.

From the root of the repository:

```
go test ./lab_four/BarrierADT
```

Add `-race` (needs cgo and gcc) to also look for data races.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
