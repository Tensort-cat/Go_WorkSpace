package main

import "fmt"

func main() {
	var i interface{} = 100
	printType(i) // int

	var j interface{} = "hello"
	printType(j) // string

	var k interface{} = 3.1415926
	printType(k) // unknown type
}

func printType(i interface{}) { // type switch
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type")
	}
}
