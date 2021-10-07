package main

import "fmt"

type PersonStringer struct {
	Name string
	Age  int
}

func (p PersonStringer) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func Stringer() {
	a := PersonStringer{"Arthur Dent", 42}
	z := PersonStringer{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z) // automatically String() function
}
