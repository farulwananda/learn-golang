package main

import "fmt"

func main() {
	name := "Farul"
	counter := 0

	increment := func() {
		name := "Wananda"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
