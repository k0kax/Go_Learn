/*

有缓存的管道，
当管道满时，再往里写数据就会阻塞
当管道空时，再往里读数据也会阻塞
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //有缓冲的channel

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子协程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子协程正在运行： 发送元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c //从c中接受数据，并赋值给num
		fmt.Println("num=", num)
	}

	fmt.Println("main结束")
}
