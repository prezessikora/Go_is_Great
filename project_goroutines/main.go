package main

import (
	"fmt"
	"time"
)

func greet(prompt string, done chan bool) {
	fmt.Println(prompt)
	done <- true
}

func slowGreet(prompt string, done chan bool) {

	time.Sleep(3 * time.Second)
	fmt.Println(prompt)
	done <- true
}

func main() {
	prices := [4]int{1, 2, 3, 4}
	fmt.Println(prices)
	dones := make([]chan bool, 4)
	dones[0] = make(chan bool)
	dones[1] = make(chan bool)
	dones[2] = make(chan bool)
	dones[3] = make(chan bool)

	go greet("Hi ..", dones[0])
	go greet("How are you? ..", dones[1])
	go slowGreet("How are you slow? ..", dones[2])
	go greet("Bye bye", dones[3])
	for _, done := range dones {
		<-done
	}
}
