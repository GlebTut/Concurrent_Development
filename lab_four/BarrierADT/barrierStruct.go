//barrierStruct.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Original author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Modified by: Gleb Tutubalin C00290944
// Modified on: 7/11/2026
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// A reusable barrier as an abstract data type (a struct with a wait method).
//
// Changes:
//   1. The mutex was replaced by an atomic counter.
//   2. The barrier is reusable: two channels (turnstiles) are used, the same
//      design as SimpleBarrier/barrier2.go.
//   3. createBarrier returns a pointer, so the barrier is never copied, and
//      the debug prints were removed.
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

const rounds = 3 // the same barrier is used this many times

// barrier is a reusable barrier for a fixed number of goroutines.
type barrier struct {
	turnstile1 chan bool    // opens when everyone has arrived
	turnstile2 chan bool    // opens when everyone has left turnstile 1
	total      int          // number of goroutines using the barrier
	count      atomic.Int32 // number of goroutines currently inside
}

// createBarrier returns a properly initialised barrier.
// N == number of goroutines
func createBarrier(N int) *barrier {
	return &barrier{
		turnstile1: make(chan bool),
		turnstile2: make(chan bool),
		total:      N,
	}
}

// wait blocks until everyone reaches this point, then lets everyone continue.
// It can be called again for the next round.
func (b *barrier) wait() {
	// Turnstile 1: the last to arrive releases the others.
	if b.count.Add(1) == int32(b.total) {
		for range b.total - 1 {
			b.turnstile1 <- true
		}
	} else {
		<-b.turnstile1
	}
	// Turnstile 2: the last to leave releases the others, so the barrier
	// is empty (count == 0) before anyone can enter it again.
	if b.count.Add(-1) == 0 {
		for range b.total - 1 {
			b.turnstile2 <- true
		}
	} else {
		<-b.turnstile2
	}
} //wait

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	for round := 1; round <= rounds; round++ {
		time.Sleep(time.Duration(rand.IntN(300)) * time.Millisecond) //wait random time amount
		fmt.Println("Part A round", round, "goroutine", Num)
		//Rendezvous here
		theBarrier.wait()
		fmt.Println("PartB round", round, "goroutine", Num)
	}
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5
	theBarrier := createBarrier(threadCount)

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, theBarrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done
}
