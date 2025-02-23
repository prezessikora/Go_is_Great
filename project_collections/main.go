package main

import (
	"fmt"
)

func dynamic_slices() {
	prices := []float64{10.98, 11.12, 1.11}
	fmt.Println(prices[1:])
	prices2 := append(prices, 9.11, 3.11, 1.11)
	fmt.Println(prices2)
	discounts := []float64{1.11, 2.11, 3.11}
	prices = append(prices, discounts...)
	fmt.Println(prices)
}

func fixed_slices() {
	prices := [4]int{1, 2, 3, 4}
	fmt.Println(prices)
	pS := prices[0:2]
	fmt.Println(pS)
	pS[0] = 99
	fmt.Println(prices)
	for index, value := range prices {
		fmt.Printf("Element at %v: %v\n", index, value)
	}
}

type smap map[string]string

func maps() {
	mymap := smap{
		"google": "http://google.com",
		"aws":    "http://amazon.com",
	}
	fmt.Println(mymap)
	fmt.Println(mymap["aws"])
	mymap["linkedin"] = "linkedin.com"
	fmt.Println(mymap)
	delete(mymap, "aws")
	fmt.Println(mymap)
	fmt.Println("=====")
	for key, value := range mymap {
		fmt.Printf("Key %v contains: %v\n", key, value)
	}
}

func create_with_make() {
	user := make([]string, 20)
	user = append(user, "Max")
	user = append(user, "Kris")
	user[0] = "Sandra"
	user[1] = "Borys"
	fmt.Println(user)
}

func main() {
	fmt.Printf("ssssss %s")

}
