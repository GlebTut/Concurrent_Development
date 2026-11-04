// Lab: Go Concurrency Essentials – Parallel Fibonacci
// Author: Gleb Tutubalin C00290944
// Based on the template code provided for the lab.
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// The template's parFib started a new goroutine for EVERY recursive call, which
// creates hundreds of millions of goroutines for fib(40) and crashes with
// "fatal error: out of memory". The fix is a threshold: for small N the
// work is cheaper than starting a goroutine, so at or below the threshold we
// call the ordinary sequential fib instead.

//--------------------------------------------
// Modified by: Gleb Tutubalin C00290944
// Modified on: 22/9/2026
// Changes:
//   - parFib now takes a threshold and uses plain fib at or below it
//   - main passes the threshold to parFib
//   - Added comments explaining the changes
// Issues: None
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
)

// fib is the ordinary sequential recursive version.
func fib(N int) int {
	if N < 2 {
		return 1
	} else {
		return fib(N-1) + fib(N-2)
	}
}

// parFib computes the same result as fib, but runs the two recursive calls in
// separate goroutines. At or below the threshold it falls back to the
// sequential fib: the work there is too small to be worth a goroutine.
func parFib(N int, threshold int) int {
	if N <= threshold {
		return fib(N)
	}

	var wg sync.WaitGroup
	var A, B int
	wg.Add(2)
	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-1, threshold)
	}(N, &A)
	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-2, threshold)
	}(N, &B)
	wg.Wait() // wait for both halves, then combine them
	return A + B
}

func main() {
	// Below this value parFib stops creating goroutines and calls fib directly.
	threshold := 25

	// Both versions must print the same number on every line.
	for i := 0; i < 10; i++ {
		Seq := fib(i * 5)
		par := parFib(i*5, threshold)
		fmt.Println(Seq, "---", par)
	}
}
