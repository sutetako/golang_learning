package main

import "fmt"

func for_is_goes_while_main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
