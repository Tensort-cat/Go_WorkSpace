// 接口是 Go 中实现多态的主要方式。通过定义接口，不同的结构体可以实现相同的方法，从而实现类似继承的多态行为
package main

import "fmt"

/*
代码解释
Speaker 是一个接口，定义了一个 Speak 方法。
Animal 结构体实现了 Speaker 接口。
Dog 结构体通过嵌入 Animal 结构体，间接实现了 Speaker 接口。
在 main 函数中，我们将 Dog 实例赋值给 Speaker 接口，并通过接口调用 Speak 方法。
*/

// 定义接口
type Speaker interface {
	Speak()
}

// 父结构体
type Animal struct {
	Name string
}

// 实现接口方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal
	Breed string
}

func main() {
	var speaker Speaker

	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	speaker = &dog
	speaker.Speak() // 通过接口调用方法
}
