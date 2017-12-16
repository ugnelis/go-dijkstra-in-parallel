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

package prio_test

import (
	prio "."
	"fmt"
)

type myType struct {
	value int
	index int // index in heap
}

func (x *myType) Less(y prio.Interface) bool { return x.value < y.(*myType).value }
func (x *myType) Index(i int)                { x.index = i }

func ExampleQueue() {
	var q prio.Queue

	a := []*myType{{2, 0}, {4, 0}, {1, 0}, {5, 0}, {3, 0}}
	for i := range a {
		q.Push(a[i])
	}

	q.Remove(a[0].index) // Locate and remove element a[0] (with value 2).
	a[1].value = 6
	q.Fix(a[1].index) // Fix heap ordering after changing value of a[1] (from 4 to 6).
	for q.Len() > 0 {
		fmt.Print(q.Pop().(*myType).value)
	}
	// Output: 1356
}
