package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now() //计时器
	var wg sync.WaitGroup
	for i := 8082; i <= 8100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("124.221.252.6:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 开启\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n用时%ds", elapsed)
}
