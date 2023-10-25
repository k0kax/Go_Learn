/*
go语言的反射
*/
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))   //获取数据类型
	fmt.Println("value:", reflect.ValueOf(arg)) //获取value
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("User is called")
	fmt.Printf("%v\n", this)
} //不用指针

func DoFileAndMethod(input interface{}) {
	//结构体
	//获取type
	fmt.Println("--------------------结构体-----------------------------")
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is", inputType.Name()) //通过type获取数据类型的string值，也就是名

	//获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputType is", inputValue)

	//结构体内部的字段
	//通过type获取里面的字段
	fmt.Println("--------------------结构体内部字段-----------------------------")
	//获取interface的reflect.Type
	for i := 0; i < inputType.NumField(); i++ { //通过Type的NumFileld，进行遍历
		field := inputType.Field(i)              //Type的Fileld获取每个数据类型
		value := inputValue.Field(i).Interface() //通过Value获取每个的值
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}

	//通过type获取结构体里面的方法，调用
	fmt.Println("--------------------结构体内部方法-----------------------------")
	for i := 0; i < inputType.NumMethod(); i++ { //获取方法数
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}

}

func main() {

	var num float64 = 3.1415926
	fmt.Println(num)
	reflectNum(num)

	user1 := User{110, "Bob", 20} //此处仅为声明一个对象
	DoFileAndMethod(user1)
}
