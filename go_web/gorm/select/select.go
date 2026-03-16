package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 可以通过tag default为字段设置默认值，插入数据时如果没有为该字段赋值，则会使用默认值
// 注意，当结构体的字段默认值是零值的时候比如 0, "", false，这些字段值将不会被保存到数据库中
type User struct {
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

	db.AutoMigrate(&User{})

	users := []User{
		{Name: "pyp", Age: 16},
		{Name: "lkw", Age: 10},
		{Name: "yh", Age: 18},
	}
	db.Create(&users)

	// 获取第一条记录 (主键升序)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	var userByFirst User
	res := db.First(&userByFirst)
	printDBRes(res, &userByFirst)

	// 获取一条记录，没有指定排序字段
	// SELECT * FROM users LIMIT 1;
	var userByTake User
	res = db.Take(&userByTake)
	printDBRes(res, &userByTake)

	// 获取最后一条记录 (主键降序)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	var userByLast User
	res = db.Last(&userByLast)
	printDBRes(res, &userByLast)

	// 根据主键检索
	// SELECT * FROM users WHERE id = 2;
	var user User
	res = db.First(&user, 2) // 2写成字符串"2"也可以
	printDBRes(res, &user)

	// SELECT * FROM users WHERE id IN (2,3);
	users = make([]User, 10)
	res = db.Find(&users, []int{2, 3})
	if res.Error != nil {
		fmt.Printf("查询失败\n")
	} else {
		fmt.Println(users)
	}

	// Get all records
	// SELECT * FROM users;
	res = db.Find(&users)

	// 条件
	// Get first matched record
	db.Where("name = ?", "lkw").First(&user)
	fmt.Println(user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// Get all matched records
	db.Where("name <> ?", "jinzhu").Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN ?", []string{"pyp", "yh"}).Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE name IN ('pyp', 'yh');

	// LIKE
	db.Where("name LIKE ?", "%y%").Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE name LIKE '%y%';

	// AND
	db.Where("name = ? AND age >= ?", "pyp", "10").Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	db.Where("updated_at > ?", "2000-01-01 00:00:00").Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", "2000-01-01 00:00:00", "2025-10-01 00:00:00").Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// Struct & Map 条件
	// Struct
	db.Where(&User{Name: "pyp", Age: 16}).First(&user)
	fmt.Println(user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]any{"name": "jinzhu", "age": 20}).Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Slice of primary keys
	db.Where([]int64{20, 21, 22}).Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	// Not 条件
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]any{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	// OR 条件
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]any{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// 排序
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []any{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

}

func printDBRes(res *gorm.DB, user *User) {
	if res.Error != nil {
		fmt.Printf("查询失败\n")
		return
	}

	fmt.Println(user)
}
