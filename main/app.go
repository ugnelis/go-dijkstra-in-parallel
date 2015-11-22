package main

import (
	"time"
	"code.google.com/p/go-priority-queue/prio"
)

func ComputePaths(source *Vertex) {
	source.minDistance = 0

	var q prio.Queue
	q.Push(&prioVertex{value:source})

	for q.Len() > 0 {
		u := q.Pop()

		for _, element := range u.(*prioVertex).value.adjacencies {
			v := &prioVertex{value:element.target}
			weight := element.weight
			distanceThroughU := u.(*prioVertex).value.minDistance + weight
			go func(val *prioVertex, distance float64) {
				if distanceThroughU < val.value.minDistance {
					val.value.minDistance = distance
					val.value.previous = u.(*prioVertex).value
					q.Push(val)
				}
			}(v, distanceThroughU)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func GetShortestPathTo(target *Vertex) []*Vertex {
	path := []*Vertex{}

	for vertex := target; vertex != nil; vertex = vertex.previous {
		path = append(path, vertex);
	}

	Reverse(path)
	return path
}

func Reverse(a []*Vertex) {
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
}

type Vertex struct {
	name        string
	adjacencies []Edge
	minDistance float64
	previous    *Vertex
}

type Edge struct {
	target *Vertex
	weight float64
}

type prioVertex struct {
	value *Vertex
	index int // index in heap
}

func (x *prioVertex) Less(y prio.Interface) bool {
	return x.value.name < y.(*prioVertex).value.name
}

func (x *prioVertex) Index(i int) {
	x.index = i
}