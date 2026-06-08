package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	for { // 先补零
		if len(num1) == len(num2) {
			break
		}
		if len(num1) < len(num2) {
			num1 = "0" + num1
		} else {
			num2 = "0" + num2
		}
	}

	res := ""
	carry := 0
	for i := len(num1) - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		y := int(num2[i] - '0')
		s := x + y + carry
		if s <= 9 {
			res = fmt.Sprintf("%d", s) + res
			carry = 0
		} else {
			res = fmt.Sprintf("%d", s-10) + res
			carry = 1
		}
	}
	if carry == 1 {
		res = "1" + res
	}

	return res
}
