package main

import "fmt"

func main() {
	// map三种创建方式
	// 方式1：上节课的方式
	// 方式2：
	myMap2 := make(map[string]string)
	myMap2["name"] = "jerry"
	myMap2["age"] = "99"
	fmt.Println(myMap2)

	// 方式3：
	myMap3 := map[string]string{
		"name": "tom",
		"age":  "88",
	}
	myMap3["sex"] = "???"
	fmt.Println(myMap3)
}
