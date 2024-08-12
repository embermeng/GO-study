package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[2:8]
	// 1. 切片可以继续切片
	slice1 := slice[1:4]
	fmt.Println(slice1)
	slice1[0] = 666
	fmt.Println(arr)
	fmt.Println(slice)

	// 2. 切片可以动态增长，append返回一个新的切片，新数组和老数组不同
	slice2 := append(slice, 88, 50)
	fmt.Println(slice2)
	// 一般是给slice追加
	slice = append(slice, 88, 50)
	fmt.Println(slice)

	// 还可以将切片追加给切片
	slice3 := []int{99, 44}
	slice = append(slice, slice3...)
	fmt.Println(slice)

	// 3. 拷贝切片
	slice4 := make([]int, 10)
	copy(slice4, slice)
	fmt.Println(slice4)
}
