package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Farul"
	sayHelloTo(firstName, "Wananda")
	sayHelloTo("Farul", "Ahmad")
}
