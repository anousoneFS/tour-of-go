package main

import (
	"fmt"
	"math"
)

type Ivalue interface {
	M()
}

type Tvalue struct {
	S string
}

func (t *Tvalue) M() {
	fmt.Println(t.S)
}

type Fvalue float64

func (f Fvalue) M() {
	fmt.Println(f)
}

func InterfaceVAlues() {
	var i Ivalue

	i = &Tvalue{"Hello"}
	describe(i)
	i.M()

	i = Fvalue(math.Pi)
	describe(i)
	i.M()
}

func describe(i Ivalue) {
	fmt.Printf("(%v, %T)\n", i, i)
}
