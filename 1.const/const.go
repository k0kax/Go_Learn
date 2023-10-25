/*
8月7日
第一天
*/
package main

import "fmt"

// 常量, 固定不可改
const pi float64 = 3.1415926

const (
	no1 = iota * 10 //使用iota可以累加
	no2
	no3
	a, b = iota + 1, iota + 2
	c, d
)

func main() {
	fmt.Println("pi=", pi, "\n", no1, no2, no3)
	fmt.Println("\n", a, b, c, d)
}
