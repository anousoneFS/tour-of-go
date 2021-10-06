package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StuctPointers() {
	v := Vertex{1, 2}
	p := &v // point to v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(&p)
	fmt.Println(&v)
}
