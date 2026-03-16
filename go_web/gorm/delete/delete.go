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

	user := User{
		ID: 5,
	}
	db.Delete(&user)
	// DELETE from users where id = 5;

	// 带额外条件的删除
	db.Where("name = ?", "pyp").Delete(&User{})
	// DELETE from users where name = "pyp";

	// 根据主键删除
	db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;

	db.Delete(&User{}, "10")
	// DELETE FROM users WHERE id = 10;

	users := make([]User, 10)
	db.Delete(&users, []int{1, 2, 3})
	// DELETE FROM users WHERE id IN (1,2,3);

	// 批量删除
	db.Where("name LIKE ?", "%jinzhu%").Delete(&User{})
	// DELETE from emails where email LIKE "%jinzhu%";

	db.Delete(&User{}, "email LIKE ?", "%jinzhu%")
	// DELETE from emails where email LIKE "%jinzhu%";

}
