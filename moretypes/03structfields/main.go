package main

import "fmt"

// Vertex is the struct
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{89, 2}
	fmt.Println(v.X)
}
