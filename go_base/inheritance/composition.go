package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal // 嵌入Animal结构体
	Breed  string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "speak!")
}

/*
代码解释
Animal 是父结构体，包含一个字段 Name 和一个方法 Speak。
Dog 是子结构体，通过嵌入 Animal 结构体，继承了 Animal 的字段和方法。
在 main 函数中，我们创建了一个 Dog 实例，并调用了 Speak 方法。
*/

func main() {
	dog := Dog{
		Animal: Animal{Name: "sana"},
		Breed:  "PYP",
	}
	fmt.Println(dog.Name)  // sana
	fmt.Println(dog.Breed) // PYP
	dog.Speak()            // sana speak!
}
