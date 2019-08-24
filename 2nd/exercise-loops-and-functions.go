package main

import (
	"fmt"
	"math"
)

const Eps = 1.0e-15

func Sqrt(x float64) (z float64) {
	z = x / 2.0
	er := (z*z - x)
	for math.Abs(er) > Eps {
		z -= er / (2 * z)
		fmt.Printf("error = %g, z = %g\n", er, z)
		er = (z*z - x)
	}
	return
}

func exercise_loops_and_functions_main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
	fmt.Println(Sqrt(5))
	fmt.Println(math.Sqrt(5))
}
