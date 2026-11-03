// Lab: Go Concurrency Essentials – Rendezvous
// Author: Gleb Tutubalin C00290944
// Based on template by: Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie)
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// This program implements a rendezvous: five goroutines each do "Part A"
// (after a random wait of 0-4 seconds) and then meet at a rendezvous point.
// No goroutine may start "Part B" until ALL five have finished "Part A".
// A mutex protects the arrival counter and a semaphore acts as a door that
// stays closed until the last goroutine arrives. The rendezvous is single
// use: every goroutine calls Wait exactly once.

//--------------------------------------------
// Original author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Modified by: Gleb Tutubalin C00290944
// Modified on: 22/9/2026
// Changes:
//   - Implemented Rendezvous type, NewRendezvous and Wait()
//   - Integrated the rendezvous into WorkWithRendezvous
//   - Added comments explaining the synchronization
// Issues: None (rendezvous implemented)
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Rendezvous is the meeting point shared by all goroutines.
type Rendezvous struct {
	mu        sync.Mutex          // protects counter (many goroutines update it)
	counter   int                 // how many goroutines have arrived so far
	total     int                 // how many goroutines must arrive before anyone continues
	turnstile *semaphore.Weighted // the "door": closed until the last goroutine arrives
}

// NewRendezvous creates a rendezvous for total goroutines.
// The turnstile starts closed: we take its only token and keep it,
// so any later Acquire has to wait until somebody calls Release.
func NewRendezvous(total int) *Rendezvous {
	ctx := context.Background()
	turnstile := semaphore.NewWeighted(1)
	_ = turnstile.Acquire(ctx, 1) // take the token: the door is now closed

	return &Rendezvous{
		turnstile: turnstile,
		total:     total,
	}
}

// Wait blocks the calling goroutine until all total goroutines have called it.
func (r *Rendezvous) Wait() {
	// Register our arrival. The mutex stops two goroutines updating the
	// counter at the same time.
	r.mu.Lock()
	r.counter++
	if r.counter == r.total {
		// We are the last to arrive: put a token in the box to open the door.
		r.turnstile.Release(1)
	}
	r.mu.Unlock() // unlock BEFORE waiting, or the last goroutine could never get in

	// Wait at the door. The first goroutine through takes the only token...
	_ = r.turnstile.Acquire(context.Background(), 1)
	// ...and immediately puts it back, which wakes up the next waiting goroutine.
	// Each goroutine passes the token on, so all of them get through.
	r.turnstile.Release(1)
}

// WorkWithRendezvous does Part A, meets the other goroutines, then does Part B.
func WorkWithRendezvous(wg *sync.WaitGroup, Num int, rendezvous *Rendezvous) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)

	rendezvous.Wait()
	//we wait here until everyone has completed Part A
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5
	rendezvous := NewRendezvous(threadCount) // one shared rendezvous for all goroutines

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, rendezvous)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done
}
