package main

import "fmt"

type FVertex struct {
	Lat, Long float64
}

func maps_main() {
	var m map[string]FVertex
	m = make(map[string]FVertex)
	m["Bell Labs"] = FVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

}
