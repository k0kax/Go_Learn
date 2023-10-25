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

	//死循环，不断获取数据
	for {
		if data, ok := <-c; ok { //ok为协程传递的值是否有数据
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main finished")

}
