package main

import "fmt"

func type_inference_main() {
	v := 42
	//v := 3.142
	//v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}
