package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "Farul Ahmad Wananda"
	//mengambil byte index 0
	var e byte = name[0]
	// konversi int ke string
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
