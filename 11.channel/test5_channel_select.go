/*
单流程下，一个协程只能监控一个channel的状态，select可以完成监控多个channel的状态
*/
package main

import (
	"fmt"
)

// 斐波那契数列
func fibonacci(c chan int, quit chan bool) {
	x, y := 0, 1

	for {
		select { //用select监控这两个协程的状态
		case c <- x: //监控c channel是否可写
			x, y = y, x+y
		case <-quit: //
			fmt.Println("quit")
			return
		}
	}
}

/*
在这个示例中，我们定义了一个fibonacci函数，它接收一个用于传递斐波那契数列值的通道c和一个用于通知协程退出的通道quit作为参数。

在fibonacci函数内部，我们初始化两个变量x和y为0和1，并使用select语句监听两个通道的状态。如果c通道可以写入，我们将当前的斐波那契数列值发送到c通道，并更新x和y的值。如果quit通道可以接收到值，表示需要退出协程，我们直接返回函数，结束协程的执行。

在main函数中，我们指定要生成斐波那契数列的长度n，并创建一个用于传递斐波那契数列值的通道c和一个用于通知退出的通道quit。接着，我们启动一个匿名协程，在该协程中通过循环从c通道中接收斐波那契数列的值，并使用fmt.Println打印到控制台。在循环执行完n次后，我们向quit通道发送一个值，通知fibonacci协程退出。最后，我们调用fibonacci函数开始生成斐波那契数列。

这样，我们就使用了fmt包、协程和select语句来实现斐波那契数列的生成，并通过通道和select语句实现了协程间的同步和退出控制
*/

func main() {

	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子协程进行：", <-c)
		}
		quit <- true
	}()

	fibonacci(c, quit)

}
