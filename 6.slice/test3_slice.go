/*

傻逼软件，草泥马

切片从0开始
slice切片的四种声明方式
*/

package main

import "fmt"

func main() {
	//第一种方式，声明slice1并初始化，默认值为1，2，3，4。长度为4
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("len=%d,slice1=%v", len(slice1), slice1)

	//第二种方式，声明slice1,长度为4，没有分配地址
	var slice2 []int
	fmt.Printf("\nslice2=%v", slice2)
	slice2 = make([]int, 3)
	slice2[2] = 10
	fmt.Printf("\nslice2=%v", slice2)

	//第三种方式，声明并分配空间
	var slice3 []int = make([]int, 7)
	fmt.Printf("\nslice3=%v\n", slice3)
	if slice3 == nil {
		fmt.Println("slice3是空切片")
	} else {
		fmt.Println("slice3非空")
	}

	//第四种方式，声明并分配空间
	slice4 := make([]int, 12)
	fmt.Printf("\nslice4=%v", slice4)

}
