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

package prio

import "testing"

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	q := BuildTestQueue(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// Quickly builds a queue of n somewhat random elements from 0..8n-1.
func BuildTestQueue(n int) Queue {
	a := make([]Interface, n)
	lfsr := uint16(0xace1) // linear feedback shift register
	for i := n - 1; i >= 0; i-- {
		bit := (lfsr>>0 ^ lfsr>>2 ^ lfsr>>3 ^ lfsr>>5) & 1
		lfsr = lfsr>>1 | bit<<15
		e := i<<3 + int(lfsr&0x7)
		a[i] = myInt(e) // Add a number from 8i..8i+7.
	}
	return New(a...)
}
