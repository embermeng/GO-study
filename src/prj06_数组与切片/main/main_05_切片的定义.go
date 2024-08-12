package main

import "fmt"

func main() {
	// 切片的定义
	// 1. 定义要给切片，用切片引用一个创建好的数组（上节课已有）

	// 2. 通过make内置函数创建切片，参数：切片类型、切片长度、切片容量
	slice1 := make([]int, 4, 20)
	fmt.Println("slice1:", slice1)
	// make底层创建一个数组，对外不可见，所以不可操作

	// 3.直接指定具体数组
	slice2 := []int{1, 3, 5}
	fmt.Println("slice2:", slice2, "长度:", len(slice2), "容量:", cap(slice2))
}
