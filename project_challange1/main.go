package main

import "fmt"

func main() {
	hobbies := [4]string{"Motorcycles", "Dogs", "Skateboarding"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	slice1 := hobbies[:2]
	slice2 := hobbies[0:2]
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 := slice1[1:3]
	fmt.Println(slice3)

	course := []string{"Learn Go", "Have Fun", "Get the Job"}
	fmt.Println(course)

	course[1] = "Be Parralel Guru!"
	fmt.Println(course)
	course = append(course, "Earn more and be better engineer...")
	fmt.Println(course)
	p1 := Product{
		title: "Kubek",
		id:    "1",
		price: 9.76,
	}
	p2 := Product{
		title: "Skateboard",
		id:    "2",
		price: 119.76,
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

type Product struct {
	title string
	id    string
	price float64
}
