/*
8月7日
第一天
*/
package main

//import 引用包需要指定目录，不在gopath路径下需要指定包
import (
	"Learn_GoLang/3.init_import/lib1" //普通导包
	"Learn_GoLang/3.init_import/lib2"

	//三种特殊导包
	mylib "Learn_GoLang/3.init_import/lib4" //重命名包，使用别名
	. "Learn_GoLang/3.init_import/lib5"     //点导入包，不用再写包名，但是可能存在函数名冲突，注意包导入要靠前

	_ "Learn_GoLang/3.init_import/lib3" //匿名导包，可能不用
)

// init函数一般执行早于main函数
func main() {
	lib1.Test()
	lib2.Test()
	mylib.TestLib4()
	TestLib5()
}
