// Lab: Go Concurrency Essentials – Mutex Example
// Author: Gleb Tutubalin C00290944
// License: MIT; see the LICENSE file in the repository root.
//
// Collaboration: Maksym Redchenko C00302240, Matvii Prokopovych C00302259
//
// This program demonstrates using a sync.Mutex to protect a shared counter
// updated by multiple goroutines. Without the mutex, the final total would
// be incorrect due to data races.

package main

import (
	"fmt"
	"sync"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup
var total int64

// adds increments the global 'total' n times while holding the mutex.
// The mutex ensures that only one goroutine updates 'total' at a time
func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex

	total = 0
	// the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println("Starting goroutine", i)
		go adds(1000, &theLock)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println("Total:", total)
}
