package main

import (
	"fmt"
	"math"
)

type Xfunc struct {
	X, Y float64
}

func Abs(v Xfunc) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Xfunc, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerFunc() {
	v := Xfunc{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
