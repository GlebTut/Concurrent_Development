# Lab 5 - Dining Philosophers

Five philosophers sit at a round table with five forks. Each philosopher thinks, picks up the
fork on their left and the fork on their right, eats, and puts both forks back. Each fork is a
buffered channel of size 1. The program ends when every philosopher has eaten 3 times.

## The deadlock and the fix

The template can deadlock: if all five philosophers pick up their left fork at the same time,
each one holds one fork and waits for the fork held by the neighbour, so nobody can ever eat.

The fix is a second buffered channel, `table`, of size 4 (the multiplex pattern). A philosopher
must put a token into `table` before picking up forks, and takes it out after putting the forks
back. At most four philosophers can try for forks at once, so with five forks at least one of
them can always get both. The philosophers can never all be stuck holding one fork each, so the
deadlock cannot happen.

## Author

Gleb Tutubalin C00290944

Based on the template by Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie).

## Requirements

- [Go](https://go.dev/dl/) (the version given in `../../go.mod`). Only the standard library is used.

## How to run

From the root of the repository:

```
go run ./lab_five/DiningPhilosophers
```

## Expected output

Lines such as `Phil:  2 was thinking` and `Phil:  2 was eating`, in a different order each run
because of the random waits. Each philosopher (0 to 4) prints 3 thinking and 3 eating lines.
The program takes roughly 20 seconds, because the waits are up to 4 seconds each, and then exits
by itself.

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
The original template is Copyright (C) 2024 Dr. Joseph Kehoe, released under the same licence.
