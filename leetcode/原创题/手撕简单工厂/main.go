package main

import "fmt"

// 接口
type AIModel interface {
	Chat()
}

// 不同实现
type GPT struct{}

func (g *GPT) Chat() {
	fmt.Println("I am GPT")
}

type DeepSeek struct{}

func (d *DeepSeek) Chat() {
	fmt.Println("I am DeepSeek")
}

// 工厂
func NewAIModel(name string) AIModel {
	switch name {
	case "GPT":
		return new(GPT)
	case "DeepSeek":
		return new(DeepSeek)
	}

	return nil
}

// 测试
func main() {
	model1 := NewAIModel("GPT")
	model1.Chat()

	model2 := NewAIModel("DeepSeek")
	model2.Chat()
}
