/*
2022年8月9号
第三天
MAP一种GO语言自带的数据结构，和slice相似，和python的字典也差不多,一对一的感觉
*/

package main

import "fmt"

func main() {
	//第一种声明方式
	var myMap1 map[string]string         //声明一个map变量，key,value均为string
	myMap1 = make(map[string]string, 10) //分配空间
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)
	//第二种
	myMap2 := make(map[int]string)
	myMap2[10] = "red"
	fmt.Println(myMap2)
	//第三种 声明加初始化
	myMap3 := map[int]string{
		1: "green",
		2: "hello",
	}
	fmt.Println(myMap3)
}
