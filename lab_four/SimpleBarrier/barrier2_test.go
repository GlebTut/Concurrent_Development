// Tests for the simple reusable barrier (lab 4).
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
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// captureOutput runs f and returns everything it printed to the screen.
// f returns false if it did not finish in time.
func captureOutput(f func() bool) (string, bool) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	out := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		out <- string(b)
	}()
	finished := f()
	os.Stdout = old
	w.Close()
	return <-out, finished
}

// All goroutines run doStuff for several rounds. The barrier is correct if
// (1) it never deadlocks, (2) in every round all "Part A" lines are printed
// before the first "PartB" line, and (3) the counter is back to 0 at the end.
func TestBarrierRounds(t *testing.T) {
	const n = 10
	var arrived atomic.Int32
	var wg sync.WaitGroup
	turnstile1 := make(chan bool)
	turnstile2 := make(chan bool)

	output, finished := captureOutput(func() bool {
		wg.Add(n)
		for i := 0; i < n; i++ {
			go doStuff(i, &arrived, n, &wg, turnstile1, turnstile2)
		}
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		select {
		case <-done:
			return true
		case <-time.After(20 * time.Second):
			return false
		}
	})
	if !finished {
		t.Fatal("timeout: the goroutines never finished (deadlock)")
	}
	if got := arrived.Load(); got != 0 {
		t.Errorf("counter is %d at the end, want 0", got)
	}

	partA := make([]int, rounds+1) // "Part A" lines seen so far, per round
	partB := make([]int, rounds+1)
	line := regexp.MustCompile(`^(Part A|PartB) round (\d+) goroutine (\d+)$`)
	for _, l := range strings.Split(strings.TrimSpace(output), "\n") {
		m := line.FindStringSubmatch(strings.TrimSpace(l))
		if m == nil {
			t.Fatalf("unexpected output line %q", l)
		}
		r, _ := strconv.Atoi(m[2])
		if m[1] == "Part A" {
			partA[r]++
		} else {
			if partA[r] != n {
				t.Errorf("round %d: a goroutine started Part B when only %d of %d had finished Part A", r, partA[r], n)
			}
			partB[r]++
		}
	}
	for r := 1; r <= rounds; r++ {
		if partA[r] != n || partB[r] != n {
			t.Errorf("round %d: %d Part A and %d Part B lines, want %d of each", r, partA[r], partB[r], n)
		}
	}
}
