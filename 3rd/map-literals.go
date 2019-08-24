package main

import "fmt"

var m = map[string]FVertex{
	"Bell Labs": FVertex{
		40.68433, -74.39967,
	},
	"Google": FVertex{
		37.42202, -122.08408,
	},
}

func map_literals_main() {
	fmt.Println(m)
}
