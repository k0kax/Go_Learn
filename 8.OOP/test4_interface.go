/*
2022年8月10号
第4天
go语言通过interface接口的方式实现多态的效果

三个条件：
1.有一个父类（有接口）
2.有子类（实现父类的全部接口方法）
3.父类的变量（指针）指向子类的具体数据变量
*/
package main

import "fmt"

// interface本质是一个指针
// 定义一个接口
// 父类
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取种类
}

// 具体的类
type Cat struct {
	color string
}

// 重写接口的方法 子类的一些方法
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}
func main() {

	var animal AnimalIF
	animal = &Cat{"yellow"}
	showAnimal(animal)

	//赋值
	dog1 := Dog{"black"}
	showAnimal(&dog1)

}
