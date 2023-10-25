/*
8月8日
第二天

数组
*/
package main

import "fmt"

// 切片是一种特殊的数组，可以不限制死数据元素的数量，可以自由添加
func print_slice(myslice []int) {
	//引用传递，传递的是切片的指针，可以改变
	fmt.Printf("myslice数据类型%T\n", myslice)
	for index, value := range myslice {
		fmt.Println("index=", index, "value=", value)
	}
	//此处改变
	myslice[2] = 555
}

func main() {
	myslice := []int{1, 222, 333, 444} //这是一种切片的声明方式
	print_slice(myslice)

	fmt.Println("")
	for index, value := range myslice {
		fmt.Println("index=", index, "value=", value)
	}

}
