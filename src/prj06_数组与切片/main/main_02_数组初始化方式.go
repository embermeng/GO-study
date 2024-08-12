package main

import "fmt"

func main() {
	// 四种初始化方式
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1=", arr1)
	var arr2 = [3]int{4, 5, 6}
	fmt.Println("arr2=", arr2)
	// [...]是规定的写法
	var arr3 = [...]int{7, 8, 9}
	fmt.Println("arr3=", arr3)
	// :前是索引，可以不按顺序写
	var arr4 = [...]int{1: 10, 0: 11, 2: 12}
	fmt.Println("arr4=", arr4)
	arr5 := [...]string{0: "tom", 1: "jerry", 2: "alice"}
	fmt.Println("arr5=", arr5)
}
