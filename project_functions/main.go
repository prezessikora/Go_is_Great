package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(&arr)
	arr = transformArray(&arr, func(i int) int {
		return i * 2
	})
	fmt.Println(&arr)
	arr = transformArray(&arr, func(i int) int {
		return i - 1
	})
	fmt.Println(&arr)
	arr = transformArray(&arr, pow)
	fmt.Println(&arr)

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr = transformArray(&arr, powerTransformer(2))
	fmt.Println(arr)

}

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
