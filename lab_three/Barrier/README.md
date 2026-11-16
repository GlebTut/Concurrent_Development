# Lab 3 - Barrier

A reusable N-goroutine barrier in Go, built from a mutex and two weighted semaphores (a two-turnstile design).

Ten goroutines each do "Part A" (after a one second sleep), then wait at the barrier.
Only when all ten have finished Part A does any of them start "Part B".

## GitHub

All the labs are in this repository: https://github.com/GlebTut/Concurrent_Development

This lab: https://github.com/GlebTut/Concurrent_Development/tree/main/lab_three/Barrier

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## How it works

A barrier makes a group of goroutines wait for each other: nobody continues until all of them have
arrived. This barrier can be used again and again (in a loop), because it has **two turnstiles**.
A turnstile is a semaphore: a goroutine does `Acquire` and then immediately `Release`, which lets
it through and passes the token on to the next waiting goroutine.

- `counter` counts the goroutines inside the barrier. A mutex protects it.
- `turnstile1` starts **closed**, `turnstile2` starts **open**.
- **Phase 1.** Each goroutine adds 1 to `counter`. The last one to arrive closes `turnstile2` and
  opens `turnstile1`. Everybody waits at `turnstile1`, so all of them pass together.
- **Phase 2.** Each goroutine subtracts 1 from `counter`. The last one to leave closes `turnstile1`
  and opens `turnstile2`. Everybody waits at `turnstile2`.

Why two turnstiles? With only one, a fast goroutine could go round the loop and enter the barrier
again while a slow goroutine is still leaving it, and the counter would be wrong. The second
turnstile makes sure everybody has left the barrier before anyone can enter it again.

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
