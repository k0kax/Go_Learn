package main

import (
	"fmt"
	"os"
)

func main() {
	//tty:pair<type :*os.File ,"E:\Code">
	tty, err := os.OpenFile("E:\\Code", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file erro", err)
	}
	fmt.Println(tty)
}
