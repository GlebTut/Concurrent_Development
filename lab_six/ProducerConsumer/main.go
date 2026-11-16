// Producer Consumer, ported to Go from the C++ template of Dr. Joseph Kehoe
// Copyright (C) 2019 Dr. Joseph Kehoe

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
// Ported to Go by: Gleb Tutubalin C00290944
// Modified on: 7/11/2026
// License: GPL-3.0; see the LICENSE file in the repository root.
//
// Main.go: producers and consumers sharing one finite buffer.
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
)

const (
	bufferSize    = 5 // the buffer is small so that producers have to wait
	numProducers  = 3
	numConsumers  = 3
	eventsPerProd = 10 // every producer creates this many events
)

// producer creates events and adds them to the buffer.
func producer(id int, theBuffer *SafeBuffer, numLoops int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numLoops; i++ {
		e := &Event{ID: i, Producer: id} //local to this goroutine
		theBuffer.Put(e)
	}
}

// consumer takes events from the buffer and consumes them.
func consumer(id int, theBuffer *SafeBuffer, numLoops int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numLoops; i++ {
		e := theBuffer.Get()
		e.Consume(id)
	}
}

func main() {
	var wg sync.WaitGroup
	aBuffer := NewSafeBuffer(bufferSize)

	// Total events produced equals total consumed, so every goroutine ends.
	total := numProducers * eventsPerProd
	perConsumer := total / numConsumers

	wg.Add(numProducers + numConsumers)
	for p := 0; p < numProducers; p++ {
		go producer(p, aBuffer, eventsPerProd, &wg)
	}
	for c := 0; c < numConsumers; c++ {
		go consumer(c, aBuffer, perConsumer, &wg)
	}
	wg.Wait() //wait here until everyone is done
	fmt.Println("All", total, "events were produced and consumed")
}
