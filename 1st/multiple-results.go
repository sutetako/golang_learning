package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func multiple_results_main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
