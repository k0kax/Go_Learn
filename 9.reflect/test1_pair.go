/*
2022年8月10号
第4天
pair 对
一个变量，包括此变量的数据类型和变量值组成一个对pair


*/

package main

import "fmt"

func main() {
	var a string
	a = "abcd"
	//pair "type:string,value:abcd"

	var allType interface{}
	allType = a
	//不同变量之间赋值pair 不变
	str, _ := allType.(string)
	fmt.Println(str)

}
