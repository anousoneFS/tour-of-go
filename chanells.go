package main

import "fmt"

// import "time"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Println(v)
	}
	c <- sum // send sum to c
}

func ChanellsWork() {
	s := []int{7, 2, 8, -9, 4, 0, 44, 23, 1, 10, 9, 7}

	c := make(chan int)
	fmt.Println("Start")
	go sum(s[:len(s)/2], c)
	fmt.Println("goroutine 1")
	go sum(s[len(s)/2:], c)
	fmt.Println("goroutine 2")
	fmt.Println("working")
	functest()
	x, y := <-c, <-c // receive from c
	fmt.Println("prepare")
	// functest()
	fmt.Println(x, y, x+y)
	fmt.Println("end leo")
}

func functest() {
	for i := 0; i < 10; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v hey\n", i)
	}
}
