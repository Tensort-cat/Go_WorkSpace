package main

import "fmt"

// 基类
type Vehicle struct {
	Brand string
}

func (v *Vehicle) Start() {
	fmt.Println(v.Brand, "started")
}

func (v Vehicle) Stop() {
	fmt.Println(v.Brand, "stopped")
}

// 派生类
type Car struct {
	Vehicle // 嵌入Vehicle
	Model   string
}

// 重写Start方法
func (c *Car) Start() {
	fmt.Println(c.Brand, c.Model, "car started")
}

func main() {
	v := Vehicle{Brand: "Toyota"}
	c := Car{
		Vehicle: Vehicle{Brand: "Honda"},
		Model:   "Civic",
	}
	v2 := Car{
		Vehicle: Vehicle{Brand: "Ford"},
	}

	v.Start()         // Toyota started
	c.Start()         // Honda Civic car started
	c.Vehicle.Start() // Honda started
	v2.Stop()
}
