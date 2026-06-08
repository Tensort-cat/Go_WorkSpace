package main

func main() {
	slice := []int{1, 2, 2, 3, 5, 6, 5, 5, 6}
	for i, v := range slice {
		println(i, v)
	}
}
