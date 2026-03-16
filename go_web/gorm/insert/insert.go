package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 可以通过tag default为字段设置默认值，插入数据时如果没有为该字段赋值，则会使用默认值
// 注意，当结构体的字段默认值是零值的时候比如 0, "", false，这些字段值将不会被保存到数据库中
type Person struct {
	gorm.Model
	Name string `gorm:"default:'anon'"`
	Age  int    `gorm:"default:16"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 根据 Person 结构体自动创建或更新表结构。
	db.AutoMigrate(&Person{})

	// 插入单条数据。
	db.Create(&Person{Name: "Alice", Age: 30})

	// 插入多条数据。
	people := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	db.Create(&people)
	for _, p := range people {
		fmt.Printf("ID: %d, name: %s, age: %d\n", p.ID, p.Name, p.Age)
	}

	// 创建记录并为指定字段赋值
	db.Select("Name", "Age").Create(&Person{Name: "pyp", Age: 18})

	// 创建记录并忽略传递给 ‘Omit’ 的字段值
	db.Omit("Name", "Age", "CreateAt").Create(&Person{
		Name: "fuck",
		Age:  10,
	})

	// 根据 map 插入数据
	db.Model(&Person{}).Create(map[string]any{
		"Name": "map",
		"Age":  20,
	})
	db.Model(&Person{}).Create([]map[string]any{
		{"Name": "lkw", "Age": "50"},
		{"Name": "lkw2", "Age": "60"},
	})
}
