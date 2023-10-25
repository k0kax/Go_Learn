/*
8月8日
第二天

数组
*/
package main

import "fmt"

// 数组 数组传参需要传递的数据严格相同，且属于值的拷贝
func test_array(myarray [4]int) {
	var array1 [10]int
	array2 := [9]int{121, 222, 333, 444, 555, 666, 777, 888, 999} //定义一个数组

	fmt.Println("\narray2数组长度：", len(array2))
	fmt.Printf("array1数据类型:%T\narray2数据类型:%T\n", array1, array2) //很奇怪的一个现象，如果不换行，显示数据格式时会显示第一个数组内的元素，可能是为了保持区分吧
	//普通打印
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}
	//使用range函数进行打印
	for index, value := range array2 {
		fmt.Println(index, value)
	}

	fmt.Println("--------------测试数组3，数组传参--------------")
	for index, value := range myarray {
		fmt.Println(index, value)
	}

}

func main() {
	myarry3 := [4]int{1111, 2222, 3333, 4444}
	test_array(myarry3)

}
