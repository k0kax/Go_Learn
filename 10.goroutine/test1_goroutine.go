/*
2022年8月11号
第5天
*/
package main

import (
	"fmt"
	"time"
)

// 子协程
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine:i = %d\n", i)
		time.Sleep(1 * time.Second) //延时一秒
	}
}

// 主协程
func main() {
	go newTask() //创建协程 执行newTask()
	i := 0

	//写个死循环，保证不退出
	for {
		i++
		fmt.Printf("main goroutine:i = %d\n", i)
		time.Sleep(1 * time.Second) //延时一秒
	}
}
