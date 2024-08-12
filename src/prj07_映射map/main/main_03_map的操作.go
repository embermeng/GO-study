package main

import "fmt"

func main() {
	myMap := make(map[string]string)
	// 增加
	myMap["name"] = "jerry"
	myMap["age"] = "99"
	myMap["sex"] = "???"
	myMap["car"] = "劳斯莱斯"
	// 修改
	myMap["name"] = "tom"
	// 删除，没有专门的清空函数，可以make一个新的空map，让原来的成为垃圾，被gc回收
	delete(myMap, "sex")
	// 查找
	value, flag := myMap["name"]
	fmt.Println(value, flag)
	// map的长度
	fmt.Println("myMap的长度:", len(myMap))
	// 遍历（只支持for-range）
	for k, v := range myMap {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println(myMap)
}
