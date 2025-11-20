package main

import "fmt"

func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})
}
