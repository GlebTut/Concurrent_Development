# Go Concurrent Development Labs

Labs for the Go Concurrent Development module (SETU Carlow). Each lab is in its own folder
with its own README that explains how to run it.

## GitHub

Repository with all the labs: https://github.com/GlebTut/Concurrent_Development

## Author

Gleb Tutubalin
C00290944

## Labs

| Lab | Folder | What it shows | Run |
|---|---|---|---|
| 1 and 2 | [go_concurrency_essentials_lab](go_concurrency_essentials_lab) | mutex, atomic, signalling, semaphore | see its README |
| 3 | [lab_three/Barrier](lab_three/Barrier) | reusable barrier (mutex and two semaphores) | `go run ./lab_three/Barrier` |
| 3 | [lab_three/Rendezvous](lab_three/Rendezvous) | rendezvous of N goroutines | `go run ./lab_three/Rendezvous` |
| 3 | [lab_three/Fibonacci](lab_three/Fibonacci) | parallel Fibonacci with a threshold | `go run ./lab_three/Fibonacci` |
| 4 | [lab_four/SimpleBarrier](lab_four/SimpleBarrier) | reusable barrier (atomic counter and channels) | `go run ./lab_four/SimpleBarrier` |
| 4 | [lab_four/BarrierADT](lab_four/BarrierADT) | reusable barrier as a data type | `go run ./lab_four/BarrierADT` |
| 5 | [lab_five/DiningPhilosophers](lab_five/DiningPhilosophers) | dining philosophers without deadlock | `go run ./lab_five/DiningPhilosophers` |
| 6 | [lab_six/ProducerConsumer](lab_six/ProducerConsumer) | producer consumer with a finite buffer | `go run ./lab_six/ProducerConsumer` |

## Requirements

[Go](https://go.dev/dl/) (the version given in [go.mod](go.mod)). Labs 1 to 3 also use
`golang.org/x/sync`; run `go mod tidy` once to download it.

## How to run the tests

Labs 3 to 6 have tests (`*_test.go` files next to the code). From the root of the repository:

```
go test ./...
```

- `go test -short ./...` skips the slow dining philosophers test (about 25 seconds otherwise).
- `go test -race ./...` also looks for data races. It needs cgo and a C compiler (gcc).
- To test one lab, give its folder, for example `go test ./lab_six/ProducerConsumer`.

## License

This repository is licensed under the [GPL-3.0 License](LICENSE).
