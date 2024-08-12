package main

import "fmt"

func main() {
	//指针
	var age int = 25
	fmt.Println(&age)

	//定义指针变量
	//ptr是一个指向int类型的指针，它存了age的地址
	var ptr *int = &age
	fmt.Println("ptr指针的地址:", &ptr)
	fmt.Println("ptr指针的值:", ptr)
	fmt.Println("ptr指针指向的数据:", *ptr)
}
