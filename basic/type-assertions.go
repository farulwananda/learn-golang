package main

import "fmt"

func random() interface{} {
	return "Farul"
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	//Safety type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Int")
	default:
		fmt.Println("Unknown")
	}
}
