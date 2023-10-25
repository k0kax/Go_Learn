package lib1

import "fmt"

// 方法大写说明对外开放，外部可以调用
func Test() {
	fmt.Println("hello,lib1")
}
func init() {
	fmt.Println("lib1引用成功")
}
