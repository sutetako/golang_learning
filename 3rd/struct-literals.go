package main

import "fmt"

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func struct_literals_main() {
	fmt.Println(v1, p, v2, v3)
}
