package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 可以通过tag default为字段设置默认值，插入数据时如果没有为该字段赋值，则会使用默认值
// 注意，当结构体的字段默认值是零值的时候比如 0, "", false，这些字段值将不会被保存到数据库中
type User struct {
	gorm.Model
	ID   int
	Name string `gorm:"default:'anon'"`
	Age  int    `gorm:"default:16"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// Save 会保留所有字段，即使字段是零值
	// 不要将 Save 和 Model一同使用, 这是 未定义的行为
	user := User{
		ID:   5,
		Name: "pdp",
		Age:  46,
	}

	db.Save(&user)

	// 更新单个列
	// 当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件，例如：
	// 根据条件更新
	db.Model(&User{}).Where("name = ?", "yh").Update("age", 20)
	// update users set age = 20 where name = "yh"

	// user有id
	db.Model(&user).Update("name", "fuck")
	// update users set name = 'fuck' where id = user.id

	// 更新多列
	// Updates 方法支持 struct 和 map[string]interface{} 参数。当使用 struct 更新时，默认情况下GORM 只会更新非零值的字段
	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(&user).Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = ?;

	// 根据 `map` 更新属性
	db.Model(&user).Updates(map[string]any{"name": "hello", "age": 18})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=?;
}
