// Tests for the parallel Fibonacci (lab 3).
// Copyright (C) 2026 Gleb Tutubalin C00290944
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// License: GPL-3.0; see the LICENSE file in the repository root.

package main

import "testing"

// fib counts 1 for N = 0 and N = 1, so fib(10) = 89.
func TestFibKnownValues(t *testing.T) {
	want := map[int]int{0: 1, 1: 1, 2: 2, 3: 3, 4: 5, 5: 8, 10: 89}
	for n, w := range want {
		if got := fib(n); got != w {
			t.Errorf("fib(%d) = %d, want %d", n, got, w)
		}
	}
}

// parFib must give the same answer as fib for every N and every threshold.
// Threshold 1 makes it create goroutines all the way down, which is the
// hardest case for the goroutine and WaitGroup code. The threshold must be at
// least 1: with 0, parFib(1, 0) would split into parFib(0) and parFib(-1),
// and fib(-1) is 1, so the answer would be 2 instead of 1.
func TestParFibMatchesFib(t *testing.T) {
	for _, threshold := range []int{1, 2, 10, 25} {
		for n := 0; n <= 20; n++ {
			if got, want := parFib(n, threshold), fib(n); got != want {
				t.Errorf("parFib(%d, %d) = %d, want %d", n, threshold, got, want)
			}
		}
	}
}
