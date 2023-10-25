/*
2023,8,16
第7 天
项目实战
*/
package main

import "fmt"

func main() {
	fmt.Println("开启服务")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
