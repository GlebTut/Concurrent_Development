# Labs 1 and 2 - Go Concurrency Essentials

Five small Go programs showing the basic synchronisation tools used in the module.
Each program lives in its own folder and has its own `main` function.

## Author

Gleb Tutubalin C00290944

## Programs

| Folder | What it demonstrates |
|---|---|
| `mutex/` | Ten goroutines increment a shared counter, protected by a `sync.Mutex` |
| `atomic/` | The same counter, using lock-free `sync/atomic` operations instead of a mutex |
| `signalling/` | Two goroutines ordered with an unbuffered channel used as a signal |
| `semaphore/` | Limiting concurrency to 5 tasks at a time with a buffered channel as a counting semaphore |
| `sem-ex/` | A worker pool limited by `golang.org/x/sync/semaphore`, computing Collatz step counts |

## Requirements

- [Go](https://go.dev/dl/) (the version given in `go.mod` of the repository root)
- Internet access the first time, so Go can download `golang.org/x/sync` (needed by `sem-ex/`)

## How to run

From the root of the repository:

```
go mod tidy
go run ./go_concurrency_essentials_lab/mutex
go run ./go_concurrency_essentials_lab/atomic
go run ./go_concurrency_essentials_lab/signalling
go run ./go_concurrency_essentials_lab/semaphore
go run ./go_concurrency_essentials_lab/sem-ex
```

## Expected output

- `mutex`: the line `Total: 10000`
- `atomic`: the number `10000`
- `signalling`: "Part A" from both goroutines (the second one after 5 seconds), then both "Part B" lines
- `semaphore`: 20 "Running task" lines, printed five at a time; takes about 8 seconds
- `sem-ex`: a list of 64 Collatz step counts

## Licence

This code is licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](../LICENSE) file in the repository root.
