package main

import "fmt"

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	q := 0
	for {
		select {
		case c <- x:
			// fmt.Println("select")
			x, y = y, x+y
		case q = <-quit:
			if q == 9 {
				fmt.Println("quit")
				return
			}
			fmt.Printf("this is %v\n", q)
		}
	}
}

func SelectWork() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			if i == 5 {
				quit <- 101
			}
		}
		quit <- 9
	}()
	fibonacciSelect(c, quit)
}
