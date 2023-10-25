package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {

	for p := range ports {
		address := fmt.Sprintf("124.221.252.6:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println(" net conn err:", err)
			results <- 0
			continue
		}
		conn.Close()
		results <- p
		//fmt.Printf("%d open", p)
		//wg.Done()
	}
}
func main() {
	ports := make(chan int, 1000) //100缓冲

	var openports []int
	results := make(chan int)

	//协程分配
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	//新协程
	go func() {
		for i := 1; i < 65535; i++ {
			ports <- i
		}
	}()

	//结果接受，主协程，先于任务分配
	for i := 1; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	//数据排序
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%dopen\n", port)
	}
}
