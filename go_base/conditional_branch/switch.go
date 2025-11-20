package main

import "fmt"

func main() {
	fmt.Println("Switch Example")

	/* 定义局部变量 */
	var grade string
	var marks int = 60

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	// type switch: 判断某个 interface 变量中实际存储的变量类型
	switch v := interface{}(marks).(type) {
	case int:
		fmt.Printf("marks is an int: %d\n", v)
	case string:
		fmt.Printf("marks is a string: %s\n", v)
	default:
		fmt.Printf("marks is of unknown type\n")
	}
}
