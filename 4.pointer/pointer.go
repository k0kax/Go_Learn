/*
8月7日
第一天
*/
package main

import "fmt"

func main() {
	a := 10
	b := 20

	swap1(a, b)
	fmt.Println("_______swap1___________\na=", a, "b=", b)

	swap2(&a, &b)
	fmt.Println("_________swap2_________\na=", a, "b=", b)
}

func swap1(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

// 使用函数利用指针改变值
func swap2(pa, pb *int) { //新建两个整形指针
	var temp int
	temp = *pa //temp=a
	*pa = *pb  //a=b
	*pb = temp //b=temp
	fmt.Println(&pa, &pb, &temp)
}
