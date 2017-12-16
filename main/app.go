package main

import (
	"sync"
	"runtime"
	"../go-priority-queue/prio"
)

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

// ComputePaths computes the minimum distance from the source
// to each vertex in the graph.
func ComputePaths(source *Vertex) {
	source.minDistance = 0

	var q prio.Queue
	q.Push(&prioVertex{value:source})

	runtime.GOMAXPROCS(1) // Number of processes is defined

	for q.Len() > 0 {
		u := q.Pop()
		var wg sync.WaitGroup

		for _, element := range u.(*prioVertex).value.adjacencies {
			wg.Add(1)
			go worker(u, element, &wg, &q) // Threads are launched.
		}
		wg.Wait()
	}
}

func worker(u interface{}, element Edge, wg *sync.WaitGroup, q *prio.Queue) {
  defer wg.Done()

  v := &prioVertex{value: element.target}
  weight := element.weight
  distanceThroughU := u.(*prioVertex).value.minDistance + weight

  if distanceThroughU < v.value.minDistance {
    v.value.minDistance = distanceThroughU
    v.value.previous = u.(*prioVertex).value
    q.Push(v)
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

// Less returns whether this element should sort before element x.
func (x *prioVertex) Less(y prio.Interface) bool {
	return x.value.name < y.(*prioVertex).value.name
}

// Index is called by the priority queue when this element is moved to index i.
func (x *prioVertex) Index(i int) {
	x.index = i
}

func GetPathStr(vertexArr []*Vertex) string {
  path_string := ""
  for _, p := range vertexArr {
    path_string += p.name
  }
  return path_string
}
