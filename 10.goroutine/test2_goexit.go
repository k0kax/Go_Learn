package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//套娃
	//匿名函数 形参空，返回值空
	go func() { //创建协程
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() //推出当前子子函数
			fmt.Println("B")
		}() //此括号为参数

		fmt.Println("A")
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine:i = %d\n", i)
		time.Sleep(1 * time.Second) //延时一秒
	}
}
