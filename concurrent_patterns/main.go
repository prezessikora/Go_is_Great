package main

import (
	"fmt"
	"time"

	"math/rand"

	"com.sikora/concurrent/fanin"
)

func boring(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v says %v", name, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c

}

func main_fanin() {
	c := fanin.FanIn(boring("Borys"), boring("Hugo"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

}

func main_basic() {
	borys := boring("Borys")
	hugo := boring("Hugo")

	for i := 0; i < 5; i++ {
		fmt.Println(<-borys)
		fmt.Println(<-hugo)
	}
	fmt.Println("Both are boring, quitting!")
}

func main() {
	//main_basic()
	main_fanin()
}
