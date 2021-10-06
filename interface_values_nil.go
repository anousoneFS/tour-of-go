package main

import "fmt"

type Inil interface {
	M()
}

type Tnil struct {
	S string
}

func (t *Tnil) M() {
	if t == nil {
		fmt.Println("<NIL>")
		return
	}
	fmt.Println(t.S)
}

func InterfaceValueNil() {
	var i Inil

	var t *Tnil
	i = t
	describenil(i)
	i.M()

	i = &Tnil{"hello"}
	describenil(i)
	i.M()
}

func describenil(i Inil) {
	fmt.Printf("(%v, %T)\n", i, i)
}
