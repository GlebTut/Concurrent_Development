// Lab: Go Concurrency Essentials – Signalling with Channels
// Author: Gleb Tutubalin C00290944
// License: MIT; see the LICENSE file in the repository root.
//
// Collaboration: Maksym Redchenko C00302240, Matvii Prokopovych C00302259
//
// This program demonstrates basic signalling between two goroutines using
// an unbuffered channel. One goroutine sends a signal, the other receives it,
// enforcing an ordering between "Part A" and "Part B" in each goroutine.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Unbuffered channel used as a simple signal / barrier.
	barrier := make(chan bool)

	// doStuffOne runs first, prints Part A, then signals and waits implicitly
	// because sending on an unbuffered channel blocks until a receiver arrives.
	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		barrier <- true // Signal and block until StuffTwo receives.
		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}
	// doStuffTwo waits 5 seconds, then receives the signal before continuing.
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here

		<-barrier // Wait for signal from doStuffOne.
		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait() //wait here until everyone (10 go routines) is done

}
