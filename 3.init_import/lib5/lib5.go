// 点导入包，不用再写包名，但是可能存在函数冲突，注意包导入要靠前
package lib5

import "fmt"

// 方法大写说明对外开放，外部可以调用
func TestLib5() {
	fmt.Println("hello,lib5")
}
func init() {
	fmt.Println("lib5 点导入包，不用再写包名，但是可能存在函数冲突，注意包导入要靠前")
}
