package main

import "fmt"

func for_continued_main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

}
