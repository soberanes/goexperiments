package main

import "fmt"

// Vertex is the struct to map
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -73.39967,
	}
	fmt.Println(m["Bell Labs"])
}
