# Lab 3 - Rendezvous

A rendezvous in Go for five goroutines, built from a mutex and a weighted semaphore (a single turnstile).

Each goroutine does "Part A" (after a random wait of 0-4 seconds), then meets the others at the
rendezvous point. Only when all five have finished Part A does any of them start "Part B".

The rendezvous is single use: every goroutine calls `Wait()` exactly once.
(For a reusable version, see the two-turnstile barrier in `../Barrier`.)

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## How it works

- A mutex protects a counter of how many goroutines have arrived.
- A semaphore with one token is used as a door. It starts closed (the token is taken in `NewRendezvous`).
- The last goroutine to arrive puts the token back, which opens the door.
- Every goroutine then takes the token and immediately puts it back, so the next waiting goroutine is woken up and all of them get through.

## Requirements

- [Go](https://go.dev/dl/) (the version given in `go.mod` of the repository root)
- Internet access the first time, so Go can download `golang.org/x/sync`

## How to run

From the root of the repository:

```
go mod tidy
go run ./lab_three/Rendezvous
```

Or from inside this folder:

```
go run rendezvous.go
```

## Expected output

After up to 4 seconds, five `Part A <n>` lines appear in random order, interleaved with the waiting,
followed by five `PartB <n>` lines. No `PartB` line ever appears before all five `Part A` lines
have been printed.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
