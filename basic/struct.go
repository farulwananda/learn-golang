package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name, "I life in", customer.Address, "My age is", customer.Age)
}

func main() {
	var farul Customer
	farul.Name = "Farul Wananda"
	farul.Address = "Indonesia"
	farul.Age = 19

	farul.sayHi("Farul")

	// fmt.Println(farul.Name)
	// fmt.Println(farul.Address)
	// fmt.Println(farul.Age)

	ahmad := Customer{
		Name:    "Ahmad",
		Address: "Indonesia",
		Age:     19,
	}
	fmt.Println(ahmad)
}
