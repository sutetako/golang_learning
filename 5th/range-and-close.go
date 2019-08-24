package main

import "fmt"

func fibonacci_range_and_close(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func range_and_close_main() {
	c := make(chan int, 10)
	go fibonacci_range_and_close(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
