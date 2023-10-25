/*
interface除了用于多态的实现，也可以作为方法的空接口（万能形参）
*/

package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("\n打印万能形参")
	fmt.Println(arg)
	//探测arg的数据类型
	_, ok := arg.(string)
	fmt.Printf("%v\n", ok)
	//fmt.Printf("%T\n", value)
	if ok {
		fmt.Println("输入是字符串")
	} else {
		fmt.Printf("输入不是字符串,是%T类型\n", arg)
	}
}

type Books struct {
	name string
}

func main() {
	book1 := Books{"The old man and Sea"}
	myFunc(book1)
	num := 109
	fmt.Printf("\n%T", num)
	myFunc(num)
	myFunc("你好")
}
