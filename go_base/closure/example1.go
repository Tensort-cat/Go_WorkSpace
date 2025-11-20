package main

func main() {
	add := adder()
	for range 3 {
		println(add())
	}
}

func adder() func() int {
	sum := 0
	return func() int {
		sum++
		return sum
	}
}
