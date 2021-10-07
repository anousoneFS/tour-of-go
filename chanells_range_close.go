package main

import (
	"fmt"
	"time"
)

// Sender
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i + 100)
		x, y = y, x+y
	}
	close(c)
}

func ChanRangeClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
