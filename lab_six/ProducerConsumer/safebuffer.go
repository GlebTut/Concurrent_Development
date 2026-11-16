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
// SafeBuffer.go: a bounded buffer that is safe to use from many goroutines.
//--------------------------------------------

package main

import "sync"

// SafeBuffer is a finite buffer of events. It uses the lecture solution with
// three semaphores:
//
//	mutex  - only one goroutine changes the buffer at a time
//	items  - counts the events in the buffer (consumers wait when it is 0)
//	spaces - counts the free places (producers wait when it is 0)
//
// Channels are used as semaphores: a receive is wait(), a send is signal().
type SafeBuffer struct {
	mutex  sync.Mutex
	buffer []*Event
	items  chan struct{}
	spaces chan struct{}
}

// NewSafeBuffer creates an empty buffer that holds at most size events.
func NewSafeBuffer(size int) *SafeBuffer {
	b := &SafeBuffer{
		buffer: make([]*Event, 0, size),
		items:  make(chan struct{}, size), // starts with value 0
		spaces: make(chan struct{}, size),
	}
	for i := 0; i < size; i++ { //spaces starts with value size
		b.spaces <- struct{}{}
	}
	return b
}

// Put adds an event. It waits while the buffer is full.
func (b *SafeBuffer) Put(e *Event) {
	<-b.spaces //Spaces.wait
	b.mutex.Lock()
	b.buffer = append(b.buffer, e)
	b.mutex.Unlock()
	b.items <- struct{}{} //Items.signal (after unlocking the mutex)
}

// Get removes and returns the oldest event. It waits while the buffer is empty.
func (b *SafeBuffer) Get() *Event {
	<-b.items //Items.wait must come before the mutex, otherwise deadlock
	b.mutex.Lock()
	e := b.buffer[0]
	b.buffer = b.buffer[1:]
	b.mutex.Unlock()
	b.spaces <- struct{}{} //Spaces.signal
	return e
}
