package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Farul"
	names[1] = "Ahmad"
	names[2] = "Wananda"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//Membuat array langsung
	var values = [4]int{
		90,
		80,
		70,
		10,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//mencari panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
