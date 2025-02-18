package main

import (
	"fmt"

	"com.sikora/binary_trees/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		Walk(t1, c1)
		close(c1)
	}()
	go func() {
		Walk(t2, c2)
		close(c2)
	}()

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		fmt.Println("v1=", v1, ":", ok1)
		fmt.Println("v2=", v2, ":", ok2)
		// chans differ in length
		if ok1 != ok2 {
			return false
		}
		// both chans exhausted
		if ok1 == false && ok2 == false {
			return true
		}
		// values differ
		if v1 != v2 {
			return false
		}
	}

}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(tree.New(1))
	fmt.Println("----")
	fmt.Println(tree.New(2))
	fmt.Println("----")
	r1_true := Same(tree.New(1), tree.New(1))
	fmt.Println(r1_true)
	r2_false := Same(tree.New(1), tree.New(2))
	fmt.Println(r2_false)
}
