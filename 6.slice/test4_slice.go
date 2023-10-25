package main

import "fmt"

func main() {

	//切片的扩容
	fmt.Println("--------------------切片的扩容-------------------")
	slice1 := make([]int, 3, 6) //声明切片的长度为3，容量为6
	//长度：左右指针的长度 容量：作指针到底层数组末尾的距离
	fmt.Printf("slice1的len=%d,cap=%d,value=%v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 1, 3) //追加2个元素1 3，长度改变
	fmt.Printf("slice1的len=%d,cap=%d,value=%v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 12121, 32121, 21212, 22121221) //追加的元素超过容量时，自动增加当前的容量的一倍
	fmt.Printf("slice1的len=%d,cap=%d,value=%v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 121212, 312121, 89898, 90899, 8778) //追加2个元素1 3，长度改变
	fmt.Printf("slice1的len=%d,cap=%d,value=%v\n", len(slice1), cap(slice1), slice1)

	fmt.Println("--------------------切片的截取-------------------")

	//切片截取的长度为所截去的长度，容量为左指针到末尾的容量
	slice2 := make([]int, 3)
	slice2 = slice1[7:] //第七个到末尾
	fmt.Printf("slice2的len=%d,cap=%d,value=%v\n", len(slice2), cap(slice2), slice2)

	slice3 := make([]int, 3)
	slice3 = slice1[0:3] // 0，1，2
	fmt.Printf("slice2的len=%d,cap=%d,value=%v\n", len(slice3), cap(slice3), slice3)

	slice4 := make([]int, 3)
	slice4 = slice1[2:4] // 2，3
	fmt.Printf("slice2的len=%d,cap=%d,value=%v\n", len(slice4), cap(slice4), slice4)

}
