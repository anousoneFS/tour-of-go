package main

import (
	"fmt"
	"math"
)

type X struct {
	X, Y float64
}

func (v X) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *X) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerReceiver() {
	v := X{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
