package main

import "fmt"

func main() {
	// 定义map：变量名 map[键类型]值类型
	var myMap map[string]string
	// 只声明map，内存没有分配空间。必须通过make函数初始化才会分配空间
	myMap = make(map[string]string, 10) // 可以存放10个键值对
	myMap["name"] = "张三"
	myMap["age"] = "20"
	fmt.Println(myMap)
}
