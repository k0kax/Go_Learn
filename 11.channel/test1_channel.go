/*
2023年8月12日
第6天

channel，管道，用于协程之间的数据通信

无缓冲的channel

*/

package main

import "fmt"

func main() {
	//定义一个channel进行协程通信
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine正在运行")

		c <- 666 //给c发送数据，写到c
	}()

	num := <-c //从c读数据赋值给num

	fmt.Println("num=", num)
	fmt.Println("main goroutine 结束")

}
