package main

import (
	"fmt"
	"math"
)

func Example4Vertices() {
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

	ComputePaths(&v0)

	for _, v := range vertices {
		fmt.Print("Distance to ", v.name, ": ", v.minDistance, "\n")
		path := GetShortestPathTo(v)
		fmt.Print("Path: ")
		for _, p := range path {
			fmt.Print(p.name)
		}
		fmt.Print("\n")
	}

	// Output:
	// Distance to A: 0
	// Path: A
	// Distance to B: 1
	// Path: AB
	// Distance to C: 2
	// Path: ABC
	// Distance to D: 3
	// Path: ABCD
}

func Example10Vertices() {
	var v0 Vertex = Vertex{name:"A"}
	var v1 Vertex = Vertex{name:"B"}
	var v2 Vertex = Vertex{name:"C"}
	var v3 Vertex = Vertex{name:"D"}
	var v4 Vertex = Vertex{name:"E"}
	var v5 Vertex = Vertex{name:"F"}
	var v6 Vertex = Vertex{name:"G"}
	var v7 Vertex = Vertex{name:"H"}
	var v8 Vertex = Vertex{name:"I"}
	var v9 Vertex = Vertex{name:"J"}

	v0.adjacencies = []Edge{Edge{&v1, 3}, Edge{&v2, 6}, Edge{&v8, 6}}
	v1.adjacencies = []Edge{Edge{&v2, 3}}
	v2.adjacencies = []Edge{Edge{&v1, 4}, Edge{&v3, 6}}
	v3.adjacencies = []Edge{Edge{&v2, 1}, Edge{&v7, 3}}
	v4.adjacencies = []Edge{Edge{&v5, 1}, Edge{&v9, 3}}
	v5.adjacencies = []Edge{Edge{&v4, 2}, Edge{&v6, 2}, Edge{&v7, 2}}
	v6.adjacencies = []Edge{Edge{&v5, 1}, Edge{&v8, 2}}
	v7.adjacencies = []Edge{Edge{&v3, 5}, Edge{&v5, 3}}
	v8.adjacencies = []Edge{Edge{&v0, 3}, Edge{&v4, 4}}
	v9.adjacencies = []Edge{Edge{&v5, 4}}

	vertices := []*Vertex{&v0, &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9}

	for _, v := range vertices {
		v.minDistance = math.MaxFloat32
	}

	ComputePaths(&v5)

	for _, v := range vertices {
		fmt.Print("Distance to ", v.name, ": ", v.minDistance, "\n")
		path := GetShortestPathTo(v)
		fmt.Print("Path: ")
		for _, p := range path {
			fmt.Print(p.name)
		}
		fmt.Print("\n")
	}

	// Output:
	// Distance to A: 7
	// Path: FGIA
	// Distance to B: 10
	// Path: FGIAB
	// Distance to C: 8
	// Path: FHDC
	// Distance to D: 7
	// Path: FHD
	// Distance to E: 2
	// Path: FE
	// Distance to F: 0
	// Path: F
	// Distance to G: 2
	// Path: FG
	// Distance to H: 2
	// Path: FH
	// Distance to I: 4
	// Path: FGI
	// Distance to J: 5
	// Path: FEJ
}