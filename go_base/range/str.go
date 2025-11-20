package main

import "fmt"

func main() {
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}
