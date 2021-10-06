package main

import "fmt"

func trace2(s string)   { fmt.Println("entering:", s) }
func untrace2(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func A2() {
	defer untrace2("a")
	trace2("a")
	// do something....
}
