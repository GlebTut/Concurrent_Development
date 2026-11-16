# Lab 4 - Simple Barrier

A reusable barrier for 10 goroutines, built from an atomic counter and two unbuffered channels (two turnstiles). Each goroutine runs 3 rounds of "Part A", waits at the barrier, then does "Part B".

## GitHub

All the labs are in this repository: https://github.com/GlebTut/Concurrent_Development

This lab: https://github.com/GlebTut/Concurrent_Development/tree/main/lab_four/SimpleBarrier

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## How it works

A barrier makes a group of goroutines wait for each other: nobody continues until all of them have
arrived. This one can be used for many rounds, so it has two phases (two turnstiles).

- `arrived` is an atomic counter (`sync/atomic`) instead of a mutex. `Add` returns the new value, so
  exactly one goroutine sees the value `max`, with no lock needed.
- `turnstile1` and `turnstile2` are unbuffered channels. A send on an unbuffered channel waits for
  a receiver, so a channel works as a signal between goroutines.
- **Phase 1.** Each goroutine adds 1 to `arrived`. The last one to arrive sends `max-1` signals on
  `turnstile1`. Every other goroutine receives one signal and continues.
- **Phase 2.** Each goroutine subtracts 1 from `arrived`. The last one to leave (the counter reaches
  0) sends `max-1` signals on `turnstile2`. This makes sure everybody has left the barrier before
  anyone can enter it again for the next round.

The last goroutine sends one signal for each waiting goroutine, and the waiting goroutines only
receive. In an earlier version one signal was passed from goroutine to goroutine ("baton"), and
the last goroutine could take that signal itself before a slow goroutine had started waiting. The
slow goroutine then waited for ever. Sending `max-1` signals avoids this.

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod`). Only the standard library is used.

## How to run

From the root of the repository:

```
go run ./lab_four/SimpleBarrier
```

## Expected output

For each round, ten `Part A round <r> goroutine <n>` lines appear, then ten `PartB round <r> goroutine <n>` lines. The order of the goroutine numbers changes from run to run, but no `PartB` line of a round appears before all ten `Part A` lines of that round.

## Tests

The test runs 10 goroutines for 3 rounds and checks that there is no deadlock, that in every round all `Part A` lines are printed before the first `PartB` line, and that the counter is back to 0 at the end.

From the root of the repository:

```
go test ./lab_four/SimpleBarrier
```

Add `-race` (needs cgo and gcc) to also look for data races.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
