/*
2022年8月9号
第三天
struct结构体
*/

package main

import "fmt"

// 声明一种新的数据类型myint，int的一个别名
type myint int

func initMyint() {
	fmt.Println("\n------------------自定义一个数据类型myint--------------------")
	var a myint = 10
	fmt.Println(a)
	fmt.Printf("type of a =%T", a)
}

// 自定义一种结构体，具有以下三种属性
type Book struct {
	id    int
	name  string
	price int
}

func printBook(book Book) {
	//传递副本
	fmt.Println("---------------------1.普通传递副本--------------------")
	fmt.Println(book)
}

func changeBook(book *Book) {
	//指针传递
	book.price = 10010
	fmt.Println("---------------------2.指针传递修改数值--------------------\n", book)
}

func main() {

	initMyint()

	//自定义一个结构体使用
	fmt.Println("\n------------------自定义一个结构体--------------------")
	var book1 Book
	book1.id = 111
	book1.name = "The old man and Sea"
	book1.price = 100
	printBook(book1)

	changeBook(&book1)
	fmt.Println(book1)

}
