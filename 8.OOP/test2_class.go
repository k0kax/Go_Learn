/*
2022年8月9号
第三天 类与对象的实现，封装
go语言实现类和对象的方法和其他面向对象的编程语言略有不同
首先，需要构造一个结构体，此处类似于类的效果
创造一个结构体的对象，
其他方法在前面引用这个结构体，便可作为此类的方法（使用此，类的所有方法便可以分开写）
*/
package main

import "fmt"

// 自定义一个结构体，实现类的效果
// 使用首字母大写，意味着此类也可以被外部的方法引用
type Hero struct {
	Name  string //此处大写与上面类似
	Ad    int
	Leval int
}

// 引用此结构体，相当于类的方法
func (this Hero) GetName() {
	//此处为拷贝
	fmt.Println("Name=", this.Name)
}

// 引用此结构体的指针，用于数据的修改，也是一种方法
func (this *Hero) SetNmae(newNname string) {
	//此处引用类的指针可以进行修改数据
	this.Name = newNname
}

func (this Hero) Show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Leval=", this.Leval)
}
func main() {
	//实现一个对象的效果
	hero := Hero{Name: "zhang3", Ad: 100, Leval: 99}
	//用对象调用类的方法
	hero.GetName()
	hero.SetNmae("Li4")
	hero.Show()
}
