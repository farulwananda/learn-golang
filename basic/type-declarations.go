package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKtpFarul NoKTP = "3213121421"
	fmt.Println(NoKtpFarul)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
