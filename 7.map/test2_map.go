package main

import "fmt"

func printMaP(CityMap map[string]string) {
	//引用传参，本身是一种指针
	fmt.Println("------------------map传参-------------------")
	for key, value := range CityMap {
		fmt.Println(key, value)
	}
}

func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "shanghai"
	cityMap["Japan"] = "tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	fmt.Println("------------------遍历-------------------")
	for key, value := range cityMap {
		fmt.Println(key, value)
	}

	//删除
	fmt.Println("------------------删除-------------------")
	delete(cityMap, "Japan")
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
	//修改
	fmt.Println("------------------修改-------------------")
	cityMap["USA"] = "DC"
	for key, value := range cityMap {
		fmt.Println(key, value)
	}

	printMaP(cityMap)
}
