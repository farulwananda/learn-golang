package main

import "fmt"

func main() {
	name := "Farul"

	switch name {
	case "Farul":
		fmt.Println("Hello Farul")
	case "Rifda":
		fmt.Println("Hello Rifda")
	default:
		fmt.Println("Kenalan Dulu")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
