package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 根据 Product 结构体自动创建或更新表结构。
	db.AutoMigrate(&Product{})

	// 创建一条测试数据。
	db.Create(&Product{Code: "D42", Price: 100})

	// 查询主键为 1 的数据，并把结果填充到 productByID 中。
	var productByID Product
	result := db.First(&productByID, 1)
	if result.Error != nil {
		fmt.Println("按主键查询失败:", result.Error)
	} else {
		fmt.Printf("按主键查询结果: %+v\n", productByID)
		fmt.Println("ID:", productByID.ID)
		fmt.Println("Code:", productByID.Code)
		fmt.Println("Price:", productByID.Price)
	}

	// 按条件查询 code = D42 的第一条数据。
	var productByCode Product
	result = db.First(&productByCode, "code = ?", "D42")
	if result.Error != nil {
		fmt.Println("按 Code 查询失败:", result.Error)
	} else {
		fmt.Printf("按 Code 查询结果: %+v\n", productByCode)
		fmt.Println("ID:", productByCode.ID)
		fmt.Println("Code:", productByCode.Code)
		fmt.Println("Price:", productByCode.Price)
	}

	// 更新查询到的数据。
	db.Model(&productByCode).Update("Price", 200)
	db.Model(&productByCode).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&productByCode).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 删除主键为 1 的数据。
	db.Delete(&productByCode, 1)
}
