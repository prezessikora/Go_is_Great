package main

import (
	"fmt"
	"time"
)

// Generalized parallel summer
// Checks no of cores and splits the work in N=no of cores go routines

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	xx()
}

func xx() {
	tick := time.NewTicker(time.Second)

	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 5)
		done <- true
	}()
	for {
		select {
		case v := <-tick.C:
			fmt.Println("tick: ", v)
		// case <-boom:
		// 	fmt.Println("BOOM!")
		// 	return
		case <-done:
			fmt.Println("BOOM!")
			return
		}
	}
}
