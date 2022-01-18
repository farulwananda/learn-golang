package main

import "fmt"

func main() {

	//membuat map secara langsung
	person := map[string]string{
		"name":    "Farul",
		"address": "Malang",
	}

	//menambahkan data map
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//function map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "eko"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	// menghapus data map
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
