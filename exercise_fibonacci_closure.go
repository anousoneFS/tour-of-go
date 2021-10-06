package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13 ...
func FibonacciV2() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		fmt.Printf("first = %v, second = %v\n", first, second)
		return ret
	}
}

func Fibonacci() func() int {
	x := -1
	y := -1
	return func() int {
		// call first time
		if x == -1 && y == -1 {
			y = 0
			return 0
			// call second time
		} else if y == 0 && x == -1 {
			y = 1
			x = 0
			return 1
		} else {
			x, y = y, x+y
			return x + y
		}
	}
}
