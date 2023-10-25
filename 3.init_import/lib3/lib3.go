package lib3

import "fmt"

// 方法大写说明对外开放，外部可以调用
func Test() {
	fmt.Println("hello,lib3")
}
func init() {
	fmt.Println("lib3引用成功")
}
