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
// Created on 30/9/2024
// Modified by: Gleb Tutubalin C00290944
// Modified on: 4/11/2026
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// Reusable barrier using an atomic counter and two channels (turnstiles).
//
// Changes:
//   1. The mutex was replaced by an atomic counter.
//   2. The barrier is reusable: a second turnstile stops goroutines from
//      re-entering the barrier before everyone has left it. The last
//      goroutine sends max-1 signals, so no signal can reach the wrong one.
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

func doStuff(goNum int, arrived *atomic.Int32, max int, wg *sync.WaitGroup, turnstile1 chan bool, turnstile2 chan bool) bool {
	for round := 1; round <= rounds; round++ {
		time.Sleep(time.Duration(rand.IntN(300)) * time.Millisecond)
		fmt.Println("Part A round", round, "goroutine", goNum)

		// Turnstile 1: wait until everyone has finished Part A.
		if arrived.Add(1) == int32(max) { //last to arrive - signal all the others to go
			for i := 0; i < max-1; i++ {
				turnstile1 <- true
			}
		} else { //not all here yet we wait until signal
			<-turnstile1
		} //end of if-else

		// Turnstile 2: wait until everyone has left turnstile 1, then reuse.
		if arrived.Add(-1) == 0 { //last to leave - signal all the others to go
			for i := 0; i < max-1; i++ {
				turnstile2 <- true
			}
		} else { //others are still inside - wait until signal
			<-turnstile2
		} //end of if-else

		fmt.Println("PartB round", round, "goroutine", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	var arrived atomic.Int32
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	turnstile1 := make(chan bool)  //unbuffered channels in place of semaphores
	turnstile2 := make(chan bool)  //second turnstile makes the barrier reusable
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, turnstile1, turnstile2)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
