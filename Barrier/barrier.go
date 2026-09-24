// Lab: Go Concurrency Essentials – Barrier Implementation
// Author: Gleb Tutubalin C00290944
// Based on template by: Dr. Joseph Kehoe (Joseph.Kehoe@setu.ie)
// License: GPL-3.0; see the LICENSE file in this directory.
//
// Collaboration: Maksym Redchenko C00302240, Matvii Prokopovych C00302259
//
// This program implements a reusable barrier using a mutex and a channel.
// N goroutines execute "Part A", then wait at the barrier. Only when all
// goroutines have reached the barrier do they proceed to "Part B".

//Barrier.go Template Code
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
// Created on: 30/9/2024
// Modified by: Gleb Tutubalin C00290944
// Modified on: 22/9/2026
// Changes:
//   - Implemented Barrier type and Wait() method
//   - Integrated barrier into worker function
//   - Added comments explaining synchronization
// Issues: None (barrier implemented)
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type Barrier struct {
	mutex      sync.Mutex
	counter    int
	total      int
	turnstile1 *semaphore.Weighted
	turnstile2 *semaphore.Weighted
}

func NewBarrier(total int) *Barrier {
	ctx := context.Background()
	turnstile1 := semaphore.NewWeighted(1)
	turnstile2 := semaphore.NewWeighted(1)

	// The first turnstile starts closed; the second starts open.
	_ = turnstile1.Acquire(ctx, 1)

	return &Barrier{
		total:      total,
		turnstile1: turnstile1,
		turnstile2: turnstile2,
	}
}

func (b *Barrier) Wait() {
	ctx := context.Background()

	b.mutex.Lock()
	b.counter++
	if b.counter == b.total {
		// Close the second turnstile and open the first one.
		_ = b.turnstile2.Acquire(ctx, 1)
		b.turnstile1.Release(1)
	}
	b.mutex.Unlock()

	// First phase: wait until every goroutine has arrived.
	_ = b.turnstile1.Acquire(ctx, 1)
	b.turnstile1.Release(1)

	b.mutex.Lock()
	b.counter--
	if b.counter == 0 {
		// Close the first turnstile and open the second one.
		_ = b.turnstile1.Acquire(ctx, 1)
		b.turnstile2.Release(1)
	}
	b.mutex.Unlock()

	// Second phase: finish the round and make the barrier reusable.
	_ = b.turnstile2.Acquire(ctx, 1)
	b.turnstile2.Release(1)
}

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, barrier *Barrier, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	barrier.Wait()
	//we wait here until everyone has completed part A
	fmt.Println("Part B", goNum)
}

func main() {
	const totalRoutines = 10

	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	barrier := NewBarrier(totalRoutines)

	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, barrier, &wg)
	}

	wg.Wait()
}
