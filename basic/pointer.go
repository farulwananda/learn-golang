package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Bondowoso", "Jawa Timur", "Indonesia"}

	// Change address1 with operator &
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Banyuwangi"

	// Change all address with operator *
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// Make new pointer
	var address4 = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	//Pointer in function
	var alamat = Address{
		City:     "Jember",
		Province: "Jawa Timur",
		Country:  "",
	}
	// var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
