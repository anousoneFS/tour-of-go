package main

import "fmt"

func main() {
	f := FibonacciV2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// what the output is ?
