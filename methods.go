package main

import (
	"fmt"
	"math"
)

type Nums struct {
	X, Y float64
}

func (v Nums) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func CallGoMethods() {
	v := Nums{3, 4}
	fmt.Println(v.Abs())
}
