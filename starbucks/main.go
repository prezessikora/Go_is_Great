package main

import (
	"fmt"
	"time"
)

type Drink struct {
	name     string
	complete chan string
}

var drinks = map[string]int{"Flat white": 2, "Filter": 4}

type Barista struct {
	name   string
	orders chan Drink
	drinks chan Drink
}

func NewBarista(name string, orders chan Drink, drinks chan Drink) Barista {
	b := Barista{name, orders, drinks}
	go b.waitForOrders()
	return b
}

func (b Barista) makeDrink(name string) {
	fmt.Println(b.name, "making", name)
	time.Sleep(time.Second * time.Duration(drinks[name]))
}

func (b Barista) waitForOrders() {
	for {
		fmt.Println(b.name, "waiting for orders")
		drink := <-b.orders
		b.makeDrink(drink.name)
		drink.complete <- drink.name
	}
}

type Customer struct {
	name   string
	orders chan Drink
	drinks chan Drink
}

func NewCustomer(name string, orders chan Drink, drinks chan Drink) Customer {
	c := Customer{name, orders, drinks}
	return c
}

func (c Customer) collectDrink(complete chan string) {
	fmt.Println(c.name, "waiting to collect drink")
	drink := <-complete
	fmt.Println(c.name, "collected a", drink)
}
func (c Customer) order(name string) {
	fmt.Println(c.name, "ordered a", name)
	complete := make(chan string)
	drink := Drink{name, complete}
	c.orders <- drink
	go c.collectDrink(complete)
}

func main() {
	orders := make(chan Drink)
	drinks := make(chan Drink)
	NewBarista("Rob", orders, drinks)
	NewBarista("Pete", orders, drinks)
	c1 := NewCustomer("Rachid", orders, drinks)
	c2 := NewCustomer("Rich", orders, drinks)

	c1.order("Filter")
	c2.order("Flat white")
	time.Sleep(time.Second * 10)
}
