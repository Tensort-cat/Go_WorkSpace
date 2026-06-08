package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instace *Singleton
	once    sync.Once
)

func GetInstace() *Singleton {
	once.Do(func() {
		instace = new(Singleton)
	})

	return instace
}

func main() {
	instance1 := GetInstace()
	instance2 := GetInstace()

	fmt.Println(instance1 == instance2)
}
