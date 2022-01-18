package main

import "fmt"

func main() {

	var name = "Farul"

	// if else statement
	if name == "Farul" {
		fmt.Println("Hello Farul")
	} else if name == "Rifda" { //kondisi lebih dari satu
		fmt.Println("Hello Rifda")
	} else {
		fmt.Println("Halo, Silahkan Perkenalkan Diri Anda")
	}

	// short statement menghitung panjang karakter
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
