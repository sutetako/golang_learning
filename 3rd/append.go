package main

import "fmt"

func append_main() {
	var s []int
	append_printSlice(s)

	s = append(s, 0)
	append_printSlice(s)

	s = append(s, 1)
	append_printSlice(s)

	s = append(s, 2, 3, 4)
	append_printSlice(s)
}

func append_printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
