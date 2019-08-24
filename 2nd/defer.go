package main

import "fmt"

func defer_main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
