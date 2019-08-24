package main

import "fmt"

func fibonacci() func() int {
	a, b, c := 0, 0, 0
	return func() int {
		if a+b+c == 0 {
			b = 1
			return 0
		}
		a = b
		b = c
		c = a + b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
