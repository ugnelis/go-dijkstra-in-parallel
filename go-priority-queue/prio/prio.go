// Copyright 2012 Stefan Nilsson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package prio provides a priority queue.
// The queue can hold elements that implement the two methods of prio.Interface.
package prio

/*
A type that implements prio.Interface can be inserted into a priority queue.

The simplest use case looks like this:

	type myInt int

	func (x myInt) Less(y prio.Interface) bool { return x < y.(myInt) }
	func (x myInt) Index(i int)                {}

To use the Remove method you need to keep track of the index of elements
in the heap, e.g. like this:

	type myType struct {
		value int
		index int // index in heap
	}

	func (x *myType) Less(y prio.Interface) bool { return x.value < y.(*myType).value }
	func (x *myType) Index(i int)                { x.index = i }
*/
type Interface interface {
	// Less returns whether this element should sort before element x.
	Less(x Interface) bool
	// Index is called by the priority queue when this element is moved to index i.
	Index(i int)
}

// Queue represents a priority queue.
// The zero value for Queue is an empty queue ready to use.
type Queue struct {
	h []Interface
}

// New returns an initialized priority queue with the given elements.
// A call of the form New(x...) uses the underlying array of x to implement
// the queue and hence might change the elements of x.
// The complexity is O(n), where n = len(x).
func New(x ...Interface) Queue {
	q := Queue{x}
	heapify(q.h)
	return q
}

// Push pushes the element x onto the queue.
// The complexity is O(log(n)), where n = q.Len().
func (q *Queue) Push(x Interface) {
	n := len(q.h)
	q.h = append(q.h, x)
	up(q.h, n) // x.Index(n) is done by up.
}

// Pop removes a minimum element (according to Less) from the queue and returns it.
// The complexity is O(log(n)), where n = q.Len().
func (q *Queue) Pop() Interface {
	h := q.h
	n := len(h) - 1
	x := h[0]
	h[0], h[n] = h[n], nil
	h = h[:n]
	if n > 0 {
		down(h, 0) // h[0].Index(0) is done by down.
	}
	q.h = h
	x.Index(-1) // for safety
	return x
}

// Peek returns, but does not remove, a minimum element (according to Less) of the queue.
func (q *Queue) Peek() Interface {
	return q.h[0]
}

// Remove removes the element at index i from the queue and returns it.
// The complexity is O(log(n)), where n = q.Len().
func (q *Queue) Remove(i int) Interface {
	h := q.h
	n := len(h) - 1
	x := h[i]
	h[i], h[n] = h[n], nil
	h = h[:n]
	if i < n {
		down(h, i) // h[i].Index(i) is done by down.
		up(h, i)
	}
	q.h = h
	x.Index(-1) // for safety
	return x
}

// Len returns the number of elements in the queue.
func (q *Queue) Len() int {
	return len(q.h)
}

// Fix reestablishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(i) followed by a Push of the new value.
// The complexity is O(log(n)) where n = q.Len().
func (q *Queue) Fix(i int) {
	up(q.h, i)
	down(q.h, i)
}

// Establishes the heap invariant in O(n) time.
func heapify(h []Interface) {
	n := len(h)
	for i := n - 1; i >= n/2; i-- {
		h[i].Index(i)
	}
	for i := n/2 - 1; i >= 0; i-- { // h[i].Index(i) is done by down.
		down(h, i)
	}
}

// Moves element at position i towards top of heap to restore invariant.
func up(h []Interface, i int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || h[parent].Less(h[i]) {
			h[i].Index(i)
			break
		}
		h[parent], h[i] = h[i], h[parent]
		h[i].Index(i)
		i = parent
	}
}

// Moves element at position i towards bottom of heap to restore invariant.
func down(h []Interface, i int) {
	for {
		n := len(h)
		left := 2*i + 1
		if left >= n || left < 0 { // left < 0 after int overflow
			h[i].Index(i)
			break
		}
		j := left
		if right := left + 1; right < n && h[right].Less(h[left]) {
			j = right
		}
		if h[i].Less(h[j]) {
			h[i].Index(i)
			break
		}
		h[i], h[j] = h[j], h[i]
		h[i].Index(i)
		i = j
	}
}
