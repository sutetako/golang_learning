package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func range_main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
