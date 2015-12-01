package main

import (
	"sync"
	"code.google.com/p/go-priority-queue/prio"
)

// ComputePaths computes the minimum distance from the source to each vertex in the graph.
func ComputePaths(source *Vertex) {
	source.minDistance = 0

	var q prio.Queue
	q.Push(&prioVertex{value:source})

	for q.Len() > 0 {
		u := q.Pop()
		var wg sync.WaitGroup

		for _, element := range u.(*prioVertex).value.adjacencies {
			v := &prioVertex{value:element.target}
			weight := element.weight
			distanceThroughU := u.(*prioVertex).value.minDistance + weight
			wg.Add(1)
			go func() { // Threads are launched.
				defer wg.Done()
				if distanceThroughU < v.value.minDistance {
					v.value.minDistance = distanceThroughU
					v.value.previous = u.(*prioVertex).value
					q.Push(v)
				}
			}()
		}
		wg.Wait()
	}
}

// GetShortestPathTo gets the shortest path to the target vertex.
func GetShortestPathTo(target *Vertex) []*Vertex {
	path := []*Vertex{}

	for vertex := target; vertex != nil; vertex = vertex.previous {
		path = append(path, vertex);
	}

	Reverse(path)
	return path
}

// Reverse returns the reverse order for data.
func Reverse(data []*Vertex) {
	for i, j := 0, len(data) - 1; i < j; i, j = i + 1, j - 1 {
		data[i], data[j] = data[j], data[i]
	}
}

type Vertex struct {
	name        string  // name of the vertex
	adjacencies []Edge  // adjacencies of the vertex
	minDistance float64 // the shortest distance from the source to this vertex in the graph
	previous    *Vertex // the previous vertex on the shortest path
}

type Edge struct {
	target *Vertex // the vertex it points to
	weight float64 // its weight
}

// prioVertex implements prio.Interface because of insertion into a priority queue.
type prioVertex struct {
	value *Vertex // its value
	index int     // index in heap
}

// Less returns whether this element should sort before element x.
func (x *prioVertex) Less(y prio.Interface) bool {
	return x.value.name < y.(*prioVertex).value.name
}

// Index is called by the priority queue when this element is moved to index i.
func (x *prioVertex) Index(i int) {
	x.index = i
}