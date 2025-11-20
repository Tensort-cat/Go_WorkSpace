package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	// 没有成功交换
	swapName1(Book1, Book2)
	fmt.Println(Book1)
	fmt.Println(Book2)

	// 使用结构体指针交换，成功交换
	swapName2(&Book1, &Book2)
	fmt.Println(Book1)
	fmt.Println(Book2)

}

// 不使用结构体指针
func swapName1(book1, book2 Books) {
	book1.title, book2.title = book2.title, book1.title
}

// 使用结构体指针
func swapName2(book1, book2 *Books) {
	book1.title, book2.title = book2.title, book1.title
}
