# Lab 3 - Parallel Fibonacci

A parallel version of the recursive Fibonacci function in Go, fixed so that it no longer crashes.

## The problem

The template's `parFib` started two new goroutines for **every** recursive call. For `fib(40)` that
is about 331 million goroutines, so the program stopped with `fatal error: out of memory`.

## The fix

A **threshold**: when `N` is at or below it, `parFib` calls the ordinary sequential `fib` instead of
starting goroutines. The small calls near the bottom of the recursion tree do very little work, so a
goroutine for each of them costs more than it saves. With a threshold of 25, `parFib(40)` starts
about 3,000 goroutines instead of 331 million.

`main` prints `fib(N)` and `parFib(N)` side by side for N = 0, 5, 10, ..., 45 (with a threshold
of 25). The two numbers on each line must be equal.

## Why a threshold of 25?

The value is a compromise, not an exact result: anything from about 20 to 30 works for N up to 45.

- **It keeps the number of goroutines small.** For N = 45, a threshold of 25 starts about 35,000
  goroutines instead of 3.7 billion (for N = 40: about 3,200). On the test machine the
  template ran out of memory at around 40 million.
- **Each goroutine still has enough work to be worth starting.** Below the threshold, `fib(25)` runs
  sequentially: about 240,000 calls, roughly 0.3 ms on the test machine. Starting a goroutine costs
  about a microsecond, so the overhead is negligible next to that.
- **Lower values are worse.** With a threshold of 15 there are about 4.4 million goroutines for
  N = 45, and creating them costs more than the parallelism gains.
- **Higher values leave work unshared.** The work is split into only a few uneven pieces, so on a
  computer with many cores some cores can sit idle.

In a quick test on a 2-core computer, thresholds from 25 to 35 all ran about twice as fast as the
sequential `fib`, and a threshold of 15 was slower than it. The best value depends on the number of
cores, so it is not fixed in this lab: `threshold` is a parameter of `parFib`.

## Author

Gleb Tutubalin C00290944

## Requirements

- [Go](https://go.dev/dl/) (the version given in `go.mod` of the repository root)

## How to run

From the root of the repository:

```
go run ./lab_three/Fibonacci
```

Or from inside this folder:

```
go run main.go
```

The run takes a few seconds.

## Expected output

Ten lines, each with the same number on both sides:

```
1 --- 1
8 --- 8
89 --- 89
987 --- 987
10946 --- 10946
121393 --- 121393
1346269 --- 1346269
14930352 --- 14930352
165580141 --- 165580141
1836311903 --- 1836311903
```

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../../LICENSE) file in the repository root.
