package main

import "fmt"

func slice_len_cap_main() {
	s := []int{2, 3, 5, 7, 11, 13}
	slice_len_cap_printSlice(s)

	s = s[:0]
	slice_len_cap_printSlice(s)

	s = s[:4]
	slice_len_cap_printSlice(s)

	s = s[2:]
	slice_len_cap_printSlice(s)
}

func slice_len_cap_printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
