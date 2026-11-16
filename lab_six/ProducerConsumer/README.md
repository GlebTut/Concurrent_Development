# Lab 6 - Producer Consumer

Producers create events and put them into a finite buffer. Consumers take events out of the
buffer and process them. The buffer holds at most 5 events, so a producer has to wait when it
is full, and a consumer has to wait when it is empty.

This is a Go port of the C++ template by Dr. Joseph Kehoe (the `Semaphore`, `Event`, `SafeBuffer`
and `main.cpp` files). The Go version uses 3 producers and 3 consumers, 10 events per producer.

## 1. Licence

GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](../../LICENSE) file in the
repository root. The original template is Copyright (C) 2019 Dr. Joseph Kehoe, released under
the same licence.

## 2. Authors

- Gleb Tutubalin C00290944 (Go port)
- Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie), original C++ template

## 3. How to install and run

Install [Go](https://go.dev/dl/). From the root of the repository:

```
go run ./lab_six/ProducerConsumer
```

## 4. Required libraries and platform

The version of Go given in `../../go.mod`. Only the standard library is used. Works on
Windows, Linux and macOS.

## 5. How it works

`SafeBuffer` uses the lecture solution for a finite buffer, with three semaphores:

- `mutex` gives exclusive access to the buffer.
- `items` counts the events in the buffer. A consumer waits on it when the buffer is empty.
- `spaces` counts the free places. A producer waits on it when the buffer is full.

`Put`: `spaces.wait`, `mutex.wait`, add, `mutex.signal`, `items.signal`.
`Get`: `items.wait`, `mutex.wait`, remove, `mutex.signal`, `spaces.signal`.

Channels play the role of the semaphores (a receive is `wait`, a send is `signal`). A consumer
must wait on `items` before it locks the mutex. If it locks the mutex first and the buffer is
empty, it waits for an item while holding the mutex, so no producer can ever add one: deadlock.

## 6. Files

- `main.go`: the `producer` and `consumer` functions and `main`.
- `safebuffer.go`: the `SafeBuffer` type with `Put` and `Get`.
- `event.go`: the `Event` type and its `Consume` method.

## 7. Expected output

Thirty lines like `Consumer 1 consumed event 7 from producer 2`, in a different order on every
run, then `All 30 events were produced and consumed`. The program ends by itself.

## 8. Tests

The tests check that events come out in order, that `Get` waits when the buffer is empty, that `Put` waits when it is full, and that with many producers and consumers every event is consumed exactly once, nothing deadlocks and the buffer never holds more than its size.

From the root of the repository:

```
go test ./lab_six/ProducerConsumer
```

Add `-race` (needs cgo and gcc) to also look for data races.

## TODO

Nothing is outstanding. The C++ files `Semaphore.h` and `Semaphore.cpp` are not needed in Go,
because channels and `sync.Mutex` replace them.