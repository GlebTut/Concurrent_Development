// Tests for the Rendezvous (lab 3).
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
	"sync/atomic"
	"testing"
	"time"
)

const testTimeout = 5 * time.Second

// waitOrFail fails the test if the goroutines do not finish in time.
// A goroutine stuck for ever means the barrier has a deadlock.
func waitOrFail(t *testing.T, wg *sync.WaitGroup) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(testTimeout):
		t.Fatal("timeout: the goroutines never finished (deadlock)")
	}
}

// Nobody may pass while one goroutine is still missing; everybody must pass
// as soon as the last one arrives.
func TestBlocksUntilAllArrive(t *testing.T) {
	const n = 5
	b := NewRendezvous(n)
	var passed atomic.Int32
	var wg sync.WaitGroup
	run := func() {
		defer wg.Done()
		b.Wait()
		passed.Add(1)
	}
	for i := 0; i < n-1; i++ {
		wg.Add(1)
		go run()
	}
	time.Sleep(200 * time.Millisecond)
	if got := passed.Load(); got != 0 {
		t.Fatalf("%d goroutines passed before the last one arrived", got)
	}
	wg.Add(1)
	go run() // the last goroutine arrives
	waitOrFail(t, &wg)
	if got := passed.Load(); got != n {
		t.Fatalf("%d goroutines passed, want %d", got, n)
	}
}
