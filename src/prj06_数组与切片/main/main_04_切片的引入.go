package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	/*
		切片构建在数组之上
		定义一个切片名字为slice，[]动态变化的数组长度不写，int是类型
		[1:3]：切出的一段片段，索引从1开始，3结束，但不包含3————[1, 3)
		切片实际是一个结构体，包含数组的指针、切片长度和切片容器三部分
	*/
	//var slice []int = arr[1:3]
	slice := arr[2:4]
	fmt.Println("原数组:", arr)
	fmt.Println("slice:", slice, "个数:", len(slice))
	// 获取切片的容量，容量可以动态变化
	fmt.Println("slice的容量为:", cap(slice))

	fmt.Printf("数组中下标为2位置的地址：%p\n", &arr[2])
	fmt.Printf("切片中下标为0位置的地址：%p\n", &slice[0])
	slice[1] = 10
	fmt.Println("slice:", slice, "arr:", arr)
}
