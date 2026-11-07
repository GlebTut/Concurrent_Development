// Dining Philosophers Template Code
// Copyright (C) 2024 Dr. Joseph Kehoe

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
// Created on 21/10/2024
// Modified by: Gleb Tutubalin C00290944
// Modified on: 7/11/2026
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// Dining philosophers with five philosophers and five forks. Each fork is a
// buffered channel of size 1.
//
// Changes:
//   1. The deadlock was removed. At most four philosophers may sit at the
//      table at once (a buffered channel of size 4, the "multiplex" pattern).
//      With five forks and at most four philosophers, at least one philosopher
//      can always pick up both forks, so the philosophers cannot all be stuck
//      holding one fork each.
//   2. Each philosopher now eats a fixed number of times (meals) instead of
//      looping forever, so wg.Done() is reached and the program ends.
//   3. Comments and README added.
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const meals = 3 // how many times each philosopher eats

// think waits a random time, then reports that philosopher index was thinking.
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// eat waits a random time, then reports that philosopher index was eating.
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// getForks picks up the left fork (index) and then the right fork (index+1).
// Sending on a fork's channel blocks while another philosopher holds it.
func getForks(index int, forks map[int]chan bool) {
	forks[index] <- true
	forks[(index+1)%5] <- true
}

// putForks puts both forks back by emptying their channels.
func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

// doPhilStuff is the life of one philosopher. table allows at most four
// philosophers to try for forks at the same time, which prevents deadlock.
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, table chan bool) {
	for i := 0; i < meals; i++ {
		think(index)
		table <- true //sit down (blocks if four are already seated)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
		<-table //leave the table
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	table := make(chan bool, philCount-1) //at most four philosophers at the table
	for N := range philCount {
		go doPhilStuff(N, &wg, forks, table)
	} //start philosophers
	wg.Wait() //wait here until everyone (5 go routines) is done

} //main
