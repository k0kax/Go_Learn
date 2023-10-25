/*
管道的关闭
*/
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //关闭协程
	}()

	//使用range迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finished")

}
