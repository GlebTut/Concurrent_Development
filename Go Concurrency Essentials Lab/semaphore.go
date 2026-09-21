// Lab: Go Concurrency Essentials – Semaphore via Channel
// Author: Gleb Tutubalin C00290944
// License: MIT; see the LICENSE file in the repository root.
//
// Collaboration: Maksym Redchenko C00302240, Matvii Prokopovych C00302259
//
// This program implements a simple counting semaphore using a buffered
// channel. The channel capacity (maxGoroutines) limits how many goroutines
// can run their "critical section" concurrently.

package main

import (
	"fmt"
	"sync"
	"time"
)

//make struct containing channel
//add init, acquire and release

// Semaphore struct with a channel to manage concurrent access.
type semaphore struct {
	theCounter chan struct{}
}

// Acquire function to acquire a semaphore slot.
func (s *semaphore) Acquire() {
	s.theCounter <- struct{}{}
}

// Release function to release a semaphore slot.
func (s *semaphore) Release() {
	<-s.theCounter
}

func main() {
	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
