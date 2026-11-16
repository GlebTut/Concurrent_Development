// Tests for the dining philosophers (lab 5).
// Copyright (C) 2026 Gleb Tutubalin C00290944
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// License: GPL-3.0; see the LICENSE file in the repository root.

package main

import (
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// This test uses the real waits of the program (up to 4 seconds each), so it
// takes about 20 to 30 seconds. "go test -short" skips it.
//
// With the table channel (at most four philosophers seated) the program must
// never deadlock, and every philosopher must eat exactly "meals" times.
func TestNoDeadlockAndEveryoneEats(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test: skipped with -short")
	}
	const philCount = 5
	forks := make(map[int]chan bool)
	for k := 0; k < philCount; k++ {
		forks[k] = make(chan bool, 1)
	}
	table := make(chan bool, philCount-1)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	out := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		out <- string(b)
	}()

	var wg sync.WaitGroup
	wg.Add(philCount)
	for i := 0; i < philCount; i++ {
		go doPhilStuff(i, &wg, forks, table)
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	finished := true
	select {
	case <-done:
	case <-time.After(2 * time.Minute):
		finished = false
	}
	os.Stdout = old
	w.Close()
	output := <-out

	if !finished {
		t.Fatal("timeout: the philosophers never finished (deadlock)")
	}
	ate := make([]int, philCount)
	for _, l := range strings.Split(output, "\n") {
		f := strings.Fields(l) // "Phil:", "2", "was", "eating"
		if len(f) == 4 && f[3] == "eating" {
			n, _ := strconv.Atoi(f[1])
			ate[n]++
		}
	}
	for i, n := range ate {
		if n != meals {
			t.Errorf("philosopher %d ate %d times, want %d", i, n, meals)
		}
	}
}
