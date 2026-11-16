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
// Event.go: the items that producers create and consumers process.
//--------------------------------------------

package main

import "fmt"

// Event is an item created by a producer and processed by a consumer.
type Event struct {
	ID       int // number of the event, counted by its producer
	Producer int // number of the producer that created it
}

// Consume processes the event. Here it just prints it.
func (e *Event) Consume(consumer int) {
	fmt.Println("Consumer", consumer, "consumed event", e.ID, "from producer", e.Producer)
}
