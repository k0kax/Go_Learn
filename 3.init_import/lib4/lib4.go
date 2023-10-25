package lib4

import "fmt"

// 方法大写说明对外开放，外部可以调用
func TestLib4() {
	fmt.Println("hello,lib4")
}
func init() {
	fmt.Println("lib4自定义包名")
}
