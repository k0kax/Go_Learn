/*
8月8日
第二天
*/
package main

import "fmt"

func f1() {
	fmt.Println("01")
}
func f2() {
	fmt.Println("02")
}
func f3() {
	fmt.Println("03   defer最后执行")
}

func defer_return() int {
	//按照栈的方式，先进后出
	defer f1()
	defer f2()
	defer f3()

	return r()
}

func r() int {
	fmt.Println("return先执行")
	return 0
}
func main() {
	//按照栈的方式，先进后出
	defer_return()
}
