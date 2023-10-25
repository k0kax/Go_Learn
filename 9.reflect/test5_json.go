/*
2022年8月11号
第5天
*/
package main

import (
	"encoding/json"
	"fmt"
)

// 写tag
type Film struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"` //切片
}

func main() {
	film1 := Film{"功夫", 2001, 10, []string{"周星驰", "张柏芝"}}

	//编码的过程，结构体 ---> json
	jsonStr, err := json.Marshal(film1)
	if err != nil {
		fmt.Println("json marshal error", err)
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)

	//解码的过程 jsonStr ---> 结构体
	film2 := Film{}
	err = json.Unmarshal(jsonStr, &film2)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", film2)
}
