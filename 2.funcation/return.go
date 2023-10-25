/*
  8月7日
  第一天

*/

package main

import "fmt"

// 单个返回值
func fool(a string, b int) int {
	fmt.Println("姓名=", a, "\n年龄=", b)

	c := b
	return c
}

// 多个返回值
func fool2(name string, age int) (string, int) {
	fmt.Println("////////////////////////fool2////////////////////////")
	name1 := name
	age1 := age
	return name1, age1
}

// 在函数处直接定义返回值类型
func fool3(name string, age int) (name1 string, age1 int) {
	fmt.Println("////////////////////////fool3/////////////////////////")
	name1 = name
	age1 = age
	return
}

// 相同类型的返回值可以简写
func fool4(name int, age int) (name1, age1 int) {
	fmt.Println("////////////////////////fool4/////////////////////////")
	name1 = name
	age1 = age
	return
}

// 相同类型的形参也可以
func fool5(name, age int) (name1, age1 int) {
	fmt.Println("////////////////////////fool5/////////////////////////")
	name1 = name
	age1 = age
	return
}

// 多个不同的形参和返回值也可以简写
func fool6(hight, weight int, name string) (hight1, weight1 int, name2 string) {
	fmt.Println("////////////////////////fool6/////////////////////////")
	hight1 = hight
	weight = weight1
	name2 = name
	return
}

func main() {
	fmt.Println("测试信息")
	c1 := fool("tom", 14)
	fmt.Println(c1)

	c2_name, c2_age := fool2("jack", 89)
	fmt.Println(c2_name, c2_age)

	c3_name, c3_age := fool3("jackson", 82323232329)
	fmt.Println(c3_name, c3_age)

	c4_name, c4_age := fool4(114514, 82323232329)
	fmt.Println(c4_name, c4_age)

	c5, c5_2 := fool5(114514, 82323232329)
	fmt.Println(c5, c5_2)

	h1, w1, n1 := fool6(12, 12, "lucy")
	fmt.Println(h1, w1, n1)
}
