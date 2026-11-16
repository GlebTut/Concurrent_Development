# Lab 3 - Barrier, Rendezvous and Parallel Fibonacci

Three programs, each in its own folder with its own README (how to run it, expected output, tests).

| Folder | What it shows |
|---|---|
| [Barrier](Barrier) | a reusable barrier built from a mutex and two semaphores |
| [Rendezvous](Rendezvous) | a rendezvous of N goroutines |
| [Fibonacci](Fibonacci) | parallel Fibonacci with a threshold, so it does not create millions of goroutines |

## GitHub

All the labs are in this repository: https://github.com/GlebTut/Concurrent_Development

This lab: https://github.com/GlebTut/Concurrent_Development/tree/main/lab_three

## Author

Gleb Tutubalin C00290944

## How to run the tests

From the root of the repository: `go test ./lab_three/...`

## Licence

GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](../LICENSE) file in the repository root.
