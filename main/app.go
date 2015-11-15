package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
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