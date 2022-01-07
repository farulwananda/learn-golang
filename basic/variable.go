package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data
	var name string

	name = "Farul Ahmad"
	fmt.Println(name)

	name = "Farul Ahmad Wananda"
	fmt.Println(name)

	// Deklarasi langsung variabel
	var friend = "Farul Wananda"
	fmt.Println(friend)

	var age = 19
	fmt.Println(age)

	// Deklarasi variabel tanpa var
	game := "Genshin"
	fmt.Println(game)

	// Deklarasi multiple variabel
	var (
		firstname = "Farul"
		lastname  = "Wananda"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
