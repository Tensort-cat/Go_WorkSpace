package main

import "fmt"

type AIModel interface {
	Chat()
}

type GPT struct{}

func (g *GPT) Chat() {
	fmt.Println("I am GPT")
}

type DeepSeek struct{}

func (d *DeepSeek) Chat() {
	fmt.Println("I am DeepSeek")
}

// 抽象工厂
type AIProviderFactory interface {
	CreateAIModel() AIModel
}

// GPT 工厂
type GPTFactory struct{}

func (gf *GPTFactory) CreateAIModel() AIModel {
	return new(GPT)
}

// deepseek 工厂
type DeepSeekFactory struct{}

func (dsf *DeepSeekFactory) CreateAIModel() AIModel {
	return new(DeepSeek)
}

// 测试
func main() {
	var factory AIProviderFactory
	factory = new(GPTFactory)

	model := factory.CreateAIModel()
	model.Chat()
}
