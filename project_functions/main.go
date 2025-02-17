package main

import (
	"fmt"
	"math"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) PrintList() {
	current := &l
	for {
		fmt.Println(current.val)
		current = current.next
		if current == nil {
			return
		}

	}
}

func (l List[T]) Length() int {
	current := &l
	len := 0
	for {
		len += 1
		current = current.next
		if current == nil {
			return len
		}
	}
}

type ListString List[string]

func main() {
	e1 := List[string]{nil, "A"}
	e2 := List[string]{&e1, "B"}
	head := List[string]{&e2, "C"}
	fmt.Println(head.Length())
	head.PrintList()
	head = List[string]{}
	head.PrintList()
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	fmt.Println(&arr)
// 	arr = transformArray(&arr, func(i int) int {
// 		return i * 2
// 	})
// 	fmt.Println(&arr)
// 	arr = transformArray(&arr, func(i int) int {
// 		return i - 1
// 	})
// 	fmt.Println(&arr)
// 	arr = transformArray(&arr, pow)
// 	fmt.Println(&arr)

// 	arr = []int{1, 2, 3, 4, 5}
// 	fmt.Println(arr)
// 	arr = transformArray(&arr, powerTransformer(2))
// 	fmt.Println(arr)

// }

func transformArray(arr *[]int, f func(int) int) []int {
	result := []int{}
	for _, value := range *arr {
		result = append(result, f(value))
	}
	return result
}

func pow(i int) int {
	return i * i
}

func powerTransformer(factor int) func(int) int {
	return func(value int) int {
		return int(math.Pow(float64(value), float64(factor)))
	}
}
