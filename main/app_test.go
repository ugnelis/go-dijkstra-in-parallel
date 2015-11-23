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

func Example26Vertices() {
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
	var v10 Vertex = Vertex{name:"K"}
	var v11 Vertex = Vertex{name:"L"}
	var v12 Vertex = Vertex{name:"M"}
	var v13 Vertex = Vertex{name:"N"}
	var v14 Vertex = Vertex{name:"O"}
	var v15 Vertex = Vertex{name:"P"}
	var v16 Vertex = Vertex{name:"Q"}
	var v17 Vertex = Vertex{name:"R"}
	var v18 Vertex = Vertex{name:"S"}
	var v19 Vertex = Vertex{name:"T"}
	var v20 Vertex = Vertex{name:"U"}
	var v21 Vertex = Vertex{name:"V"}
	var v22 Vertex = Vertex{name:"W"}
	var v23 Vertex = Vertex{name:"X"}
	var v24 Vertex = Vertex{name:"Y"}
	var v25 Vertex = Vertex{name:"Z"}

	v0.adjacencies = []Edge{Edge{&v20, 3}, Edge{&v21, 3}}
	v1.adjacencies = []Edge{Edge{&v0, 3}}
	v2.adjacencies = []Edge{Edge{&v4, 2}}
	v3.adjacencies = []Edge{Edge{&v1, 4}, Edge{&v2, 2}}
	v4.adjacencies = []Edge{Edge{&v3, 2}, Edge{&v5, 1}}
	v5.adjacencies = []Edge{Edge{&v4, 2}}
	v6.adjacencies = []Edge{Edge{&v5, 3}}
	v7.adjacencies = []Edge{Edge{&v6, 5}, Edge{&v8, 3}}
	v8.adjacencies = []Edge{Edge{&v12, 4}, Edge{&v13, 4}}
	v9.adjacencies = []Edge{Edge{&v10, 3}}
	v10.adjacencies = []Edge{Edge{&v25, 3}}
	v11.adjacencies = []Edge{Edge{&v9, 3}}
	v12.adjacencies = []Edge{Edge{&v8, 1}, Edge{&v11, 5}}
	v13.adjacencies = []Edge{Edge{&v7, 4}, Edge{&v8, 3}, Edge{&v14, 3}, Edge{&v24, 3}}
	v14.adjacencies = []Edge{Edge{&v7, 4}, Edge{&v13, 1}, Edge{&v17, 3}}
	v15.adjacencies = []Edge{Edge{&v7, 3}}
	v16.adjacencies = []Edge{Edge{&v15, 2}}
	v17.adjacencies = []Edge{Edge{&v16, 1}}
	v18.adjacencies = []Edge{Edge{&v19, 2}, Edge{&v20, 2}}
	v19.adjacencies = []Edge{Edge{&v18, 1}, Edge{&v22, 3}, Edge{&v24, 5}}
	v20.adjacencies = []Edge{Edge{&v0, 1}}
	v21.adjacencies = []Edge{Edge{&v18, 4}}
	v22.adjacencies = []Edge{Edge{&v19, 3}}
	v23.adjacencies = []Edge{Edge{&v19, 2}, Edge{&v22, 4}}
	v24.adjacencies = []Edge{Edge{&v13, 2}, Edge{&v23, 3}, Edge{&v25, 1}}
	v25.adjacencies = []Edge{Edge{&v11, 3}, Edge{&v24, 1}}

	vertices := []*Vertex{
		&v0, &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, &v10, &v11, &v12,
		&v13, &v14, &v15, &v16, &v17, &v18, &v19, &v20, &v21, &v22, &v23, &v24, &v25}

	for _, v := range vertices {
		v.minDistance = math.MaxFloat32
	}

	ComputePaths(&v18)

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
	// Distance to A: 3
	// Path: SUA
	// Distance to B: 29
	// Path: STYNHGFEDB
	// Distance to C: 27
	// Path: STYNHGFEDC
	// Distance to D: 25
	// Path: STYNHGFED
	// Distance to E: 23
	// Path: STYNHGFE
	// Distance to F: 21
	// Path: STYNHGF
	// Distance to G: 18
	// Path: STYNHG
	// Distance to H: 13
	// Path: STYNH
	// Distance to I: 12
	// Path: STYNI
	// Distance to J: 14
	// Path: STYZLJ
	// Distance to K: 17
	// Path: STYZLJK
	// Distance to L: 11
	// Path: STYZL
	// Distance to M: 16
	// Path: STYNIM
	// Distance to N: 9
	// Path: STYN
	// Distance to O: 12
	// Path: STYNO
	// Distance to P: 18
	// Path: STYNORQP
	// Distance to Q: 16
	// Path: STYNORQ
	// Distance to R: 15
	// Path: STYNOR
	// Distance to S: 0
	// Path: S
	// Distance to T: 2
	// Path: ST
	// Distance to U: 2
	// Path: SU
	// Distance to V: 6
	// Path: SUAV
	// Distance to W: 5
	// Path: STW
	// Distance to X: 10
	// Path: STYX
	// Distance to Y: 7
	// Path: STY
	// Distance to Z: 8
	// Path: STYZ
}