package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func interface_are_satisified_implicity_main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	/*
		var j I
		describe(j)
		j.M()
	*/

	var j interface{}
	describe_empty(j)

	j = 42
	describe_empty(j)

	j = "hello"
	describe_empty(j)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describe_empty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
