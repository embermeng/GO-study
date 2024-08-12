package main

import "fmt"

func main() {
	str := "golang"
	fmt.Println(len(str))
	// new分配内存，new函数返回一个对应类型的指针
	num := new(int)
	fmt.Printf("num的类型：%T，值：%v, 地址：%v，其指向的值：%v\n", num, num, &num, *num)
}
