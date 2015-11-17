package main

import (
	"fmt"
	"math"
	"code.google.com/p/go-priority-queue/prio"
)

func main() {
	var v0 Vertex = Vertex{name:"A"}
	var v1 Vertex = Vertex{name:"B"}
	var v2 Vertex = Vertex{name:"C"}
	var v3 Vertex = Vertex{name:"D"}

	v0.adjacencies = []Edge{Edge{&v1, 1}, Edge{&v2, 3}}
	v1.adjacencies = []Edge{Edge{&v0, 2}, Edge{&v2, 1}}
	v2.adjacencies = []Edge{Edge{&v3, 1}}
	v3.adjacencies = []Edge{Edge{&v0, 3}}

	vertices := []*Vertex{&v0, &v1, &v2, &v3}

	for _, v := range vertices {
		v.minDistance = math.MaxFloat32
	}

	ComputePaths(&v3)
	fmt.Println(v3.name)
}

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

			if distanceThroughU < v.value.minDistance {
				v.value.minDistance = distanceThroughU
				v.value.previous = u.(*prioVertex).value
				q.Push(v)
			}
		}
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