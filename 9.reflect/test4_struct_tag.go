/*
tag 一种标签，用于数据的说明，配合json使用效果更好
*/

package main

import (
	"fmt"
	"reflect"
)

type resum struct {
	Name string `info:"name" doc:"my name"` //info的key为name
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() //结构体的全部元素

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", taginfo, "doc:", tagdoc)
	}
}

func main() {
	var re resum

	findTag(&re)
	/*
			findTag(&re)中的&re是将结构体实例re的地址传递给findTag函数。

			在findTag函数的定义中，参数str的类型是interface{}，也就是可以接受任意类型的参数。然而，为了在函数内部修改结构体的值，我们需要传递结构体实例的指针。

			通过使用&re，我们获取了结构体实例re的地址，并将该地址传递给findTag函数。在函数内部，通过reflect.TypeOf(str).Elem()获取到结构体的类型，并使用指针类型的Elem()方法获取到实际的结构体类型。

			这样，我们就可以遍历结构体的字段并获取标签信息了。

			总结起来，findTag(&re)中的&re是为了传递结构体实例的指针，以便在函数内部修改结构体的值。

		// Elem方法返回类型的元素类型。
		// 如果类型的种类不是数组（Array）、通道（Chan）、映射（Map）、指针（Pointer）或切片（Slice），则会引发 panic 异常。
	*/

	//re2 := resum{"tom", "man"}
	//findTag(re2)

}
