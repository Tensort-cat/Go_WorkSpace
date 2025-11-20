package main

import "fmt"

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")
}
