package main

import "fmt"

func main() {
	const firstname string = "Farul"
	const lastname = "Wananda"
	//Walaupun const tidak dipakai tidak akan error berbeda dengan variabel
	const value = 2002

	fmt.Println(firstname)
	fmt.Println(lastname)

	//Deklatasi multiple constant
	const (
		depan    string = "Farul"
		belakang        = "Ahmad"
	)

	fmt.Println(depan)
	fmt.Println(belakang)
}
