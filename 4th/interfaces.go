package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func interface_main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	// a = v
}
