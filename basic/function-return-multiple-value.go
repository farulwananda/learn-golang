package main

import "fmt"

func getFullName() (string, string, string) {
	return "Farul", "Ahmad", "Wananda"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	fmt.Println(lastName)

}
