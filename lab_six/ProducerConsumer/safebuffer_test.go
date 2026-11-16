// Tests for the producer consumer buffer (lab 6).
// Copyright (C) 2026 Gleb Tutubalin C00290944
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// License: GPL-3.0; see the LICENSE file in the repository root.

package main

import (
	"sync"
	"testing"
	"time"
)

// isBlocked reports whether ch is still empty after a short wait.
func isBlocked(ch chan *Event) bool {
	select {
	case <-ch:
		return false
	case <-time.After(150 * time.Millisecond):
		return true
	}
}

// Events come out in the order they went in (one producer, one consumer).
func TestFirstInFirstOut(t *testing.T) {
	b := NewSafeBuffer(3)
	for i := 0; i < 3; i++ {
		b.Put(&Event{ID: i})
	}
	for i := 0; i < 3; i++ {
		if e := b.Get(); e.ID != i {
			t.Errorf("Get returned event %d, want %d", e.ID, i)
		}
	}
}

// A consumer must wait while the buffer is empty, and wake up on Put.
func TestGetWaitsWhenEmpty(t *testing.T) {
	b := NewSafeBuffer(2)
	got := make(chan *Event, 1)
	go func() { got <- b.Get() }()
	if !isBlocked(got) {
		t.Fatal("Get returned from an empty buffer")
	}
	go b.Put(&Event{ID: 7}) // in a goroutine, so a broken buffer cannot hang the test
	select {
	case e := <-got:
		if e.ID != 7 {
			t.Errorf("Get returned event %d, want 7", e.ID)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Get did not wake up after Put")
	}
}

// A producer must wait while the buffer is full, and wake up when a consumer
// makes space.
func TestPutWaitsWhenFull(t *testing.T) {
	b := NewSafeBuffer(2)
	b.Put(&Event{ID: 1})
	b.Put(&Event{ID: 2})
	added := make(chan *Event, 1)
	go func() {
		b.Put(&Event{ID: 3})
		added <- nil
	}()
	if !isBlocked(added) {
		t.Fatal("Put added a third event to a buffer of size 2")
	}
	go b.Get()
	select {
	case <-added:
	case <-time.After(2 * time.Second):
		t.Fatal("Put did not wake up after Get")
	}
}

// Many producers and consumers: every event is consumed exactly once, nothing
// deadlocks, and the buffer never holds more than its size.
func TestManyProducersAndConsumers(t *testing.T) {
	const size, producers, consumers, perProducer = 3, 4, 4, 50
	b := NewSafeBuffer(size)

	var mu sync.Mutex
	seen := make(map[[2]int]int)
	maxLen := 0

	stop := make(chan struct{})
	monitorDone := make(chan struct{})
	go func() { // watches the length of the buffer all the time
		defer close(monitorDone)
		for {
			select {
			case <-stop:
				return
			default:
			}
			b.mutex.Lock()
			if l := len(b.buffer); l > maxLen {
				maxLen = l
			}
			b.mutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for p := 0; p < producers; p++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perProducer; i++ {
				b.Put(&Event{ID: i, Producer: p})
			}
		}()
	}
	for c := 0; c < consumers; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < producers*perProducer/consumers; i++ {
				e := b.Get()
				mu.Lock()
				seen[[2]int{e.Producer, e.ID}]++
				mu.Unlock()
			}
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("timeout: producers and consumers never finished (deadlock)")
	}
	close(stop)
	<-monitorDone

	if maxLen > size {
		t.Errorf("the buffer held %d events, more than its size %d", maxLen, size)
	}
	if len(seen) != producers*perProducer {
		t.Errorf("%d different events consumed, want %d", len(seen), producers*perProducer)
	}
	for e, n := range seen {
		if n != 1 {
			t.Errorf("event %v was consumed %d times, want 1", e, n)
		}
	}
}
