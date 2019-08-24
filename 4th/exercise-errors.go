package main

import (
	"fmt"
	"math"
)

const Eps = 1.0e-15

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2.0
	er := (z*z - x)
	for math.Abs(er) > Eps {
		z -= er / (2 * z)
		fmt.Printf("error = %g, z = %g\n", er, z)
		er = (z*z - x)
	}
	return z, nil
}

func exercise_errors_main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
