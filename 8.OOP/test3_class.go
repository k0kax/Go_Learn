/**
go语言的继承真难看

**/

package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Eat() {
	fmt.Println("Huamn:", this.Name, " is Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human:", this.Name, " is Walk()")
}

// 类的继承可以通过，结构体的再次引用进行实现
type SuperMan struct {
	Human
	Level int
}

// 重写父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan:", this.Name, " is  not Eat()")
}

func (this *SuperMan) Show() {
	fmt.Println("\nSuperMan show\nname:", this.Name, "\nsex:", this.Sex, "\nlevel:", this.Level)
}
func main() {
	h1 := Human{"zahng3", "man"}
	h1.Eat()
	h1.Walk()

	s1 := SuperMan{Human{"li4", "women"}, 15} //此种写法很复杂，再次采用另一种
	s1.Eat()
	s1.Walk()

	//这个方法也不简单
	var s2 SuperMan
	s2.Name = "wang5"
	s2.Sex = "man"
	s2.Level = 99

	s2.Show()

}
